package interpreter

import (
	"errors"
	"fmt"

	"github.com/wreulicke/tincaml/ast"
)

type Env map[string]interface{}

type Closure struct {
	Env      Env
	Function *ast.FunctionNode
}

func (c *Closure) String() string {
	return fmt.Sprintf("Closure(Function: %v)", c.Function)
}

func NewEnv() Env {
	return Env{}
}

func (env Env) Clone() Env {
	newEnv := Env{}
	for k, v := range env {
		newEnv[k] = v
	}
	return newEnv
}

func Evaluate(tree *ast.Tree) error {
	env := NewEnv()
	for _, n := range tree.Body {
		value, err := EvaluateExpression(n, env)
		if err != nil {
			return err
		}
		fmt.Println(value)
	}
	return nil
}

func EvaluateExpression(v ast.AST, env Env) (interface{}, error) { // TODO 値の持ち方を内部表現に変えたい
	switch node := v.(type) {
	case *ast.AdditionExpressionNode:
		return evaluateAddition(node, env)
	case *ast.MultiplicativeExpressionNode:
		return evaluateMultiplicative(node, env)
	case *ast.EqualityExpressionNode:
		return evaluateEqual(node, env)
	case *ast.NotEqualityExpressionNode:
		return evaluateNotEqual(node, env)
	case *ast.NotExpressionNode:
		return evaluateNot(node, env)
	case *ast.BooleanNode:
		return node.Value, nil
	case *ast.NumberNode:
		return node.Value, nil
	case *ast.StringNode:
		return node.Value, nil
	case *ast.RelationalExpressionNode:
		return evaluateRelational(node, env)
	case *ast.FunctionCall:
		return evaluateFunctionCall(node, env)
	case *ast.FunctionNode:
		c := &Closure{
			Function: node,
		}
		env[string(node.ID)] = c
		c.Env = env.Clone()
		return c, nil
	case *ast.NegativeNode:
		return evaluateNegative(node, env)
	case *ast.Identifier:
		v, ok := env[string(node.ID)]
		if !ok {
			return nil, fmt.Errorf("'%s' is not found", string(node.ID))
		}
		return v, nil
	case *ast.IfExpressionNode:
		return evaluateIf(node, env)
	case *ast.AssignmentExpressionNode:
		return evaluateAssignment(node, env)
	case *ast.EmptyExpressionNode:
		return nil, nil
	default:
		return nil, errors.New("Unexpected condition. cannot evaluate this node")
	}
}

func evaluateNegative(node *ast.NegativeNode, env Env) (interface{}, error) {
	v, err := EvaluateExpression(node.Node, env)
	if err != nil {
		return nil, err
	}
	if f, ok := v.(float64); ok {
		return -f, nil
	}
	return nil, errors.New("node is not number")
}

func evaluateAssignment(node *ast.AssignmentExpressionNode, env Env) (interface{}, error) {
	v, err := EvaluateExpression(node.Initializer, env)
	if err != nil {
		return nil, err
	}
	env[string(node.ID)] = v
	return v, nil
}

func evaluateRelational(node *ast.RelationalExpressionNode, env Env) (interface{}, error) {
	l, err := EvaluateExpression(node.Left, env)
	if err != nil {
		return nil, err
	}
	r, err := EvaluateExpression(node.Right, env)
	if err != nil {
		return nil, err
	}
	if lv, ok := l.(float64); ok {
		if rv, ok := r.(float64); ok {
			switch node.Operator {
			case ast.LESS:
				return lv < rv, nil
			case ast.GREATER:
				return lv > rv, nil
			case ast.LESS_EQUAL:
				return lv >= rv, nil
			case ast.GREATER_EQUAL:
				return lv >= rv, nil
			}
		}
		return nil, errors.New("right value is not number")
	}
	return nil, errors.New("left value is not number")
}

func evaluateBodies(exprs []ast.AST, env Env) (interface{}, error) {
	var val interface{}
	var err error
	for _, expr := range exprs {
		val, err = EvaluateExpression(expr, env)
		if err != nil {
			return nil, err
		}
	}
	return val, err
}

func evaluateIf(node *ast.IfExpressionNode, env Env) (interface{}, error) {
	b, err := EvaluateExpression(node.Cond, env)
	if err != nil {
		return nil, err
	}
	if v, ok := b.(bool); ok {
		if v {
			return evaluateBodies(node.Then, env)
		}
		return evaluateBodies(node.Else, env)
	}
	return nil, errors.New("Condition is not boolean")
}

func evaluateClosure(node *Closure, args []ast.AST, parentEnv Env) (interface{}, error) {
	if len(node.Function.Params) != len(args) {
		return nil, fmt.Errorf("size of parameters is expected to %d. but size of arguments is %d", len(node.Function.Params), len(args))
	}
	functionEnv := parentEnv.Clone()
	for k, v := range node.Env {
		functionEnv[k] = v
	}
	for k, v := range args {
		id := node.Function.Params[k]
		arg, err := EvaluateExpression(v, parentEnv)
		if err != nil {
			return nil, err
		}
		functionEnv[string(id.ID)] = arg
	}
	r, err := evaluateBodies(node.Function.Body, functionEnv)
	if err != nil {
		return nil, err
	}
	if f, ok := r.(*ast.FunctionNode); ok {
		return &Closure{
			Env:      functionEnv,
			Function: f,
		}, nil
	}
	return r, nil
}

func evaluateFunctionCall(node *ast.FunctionCall, parentEnv Env) (interface{}, error) {
	v, error := EvaluateExpression(node.Function, parentEnv)
	if error != nil {
		return nil, error
	}
	f, ok := v.(*Closure)
	if !ok {
		return nil, fmt.Errorf("node is not function. found: %v", v)
	}
	return evaluateClosure(f, node.Args, parentEnv)
}

func evaluateMultiplicative(node *ast.MultiplicativeExpressionNode, env Env) (interface{}, error) {
	l, err := EvaluateExpression(node.Left, env)
	if err != nil {
		return nil, err
	}
	r, err := EvaluateExpression(node.Right, env)
	if err != nil {
		return nil, err
	}
	if lv, ok := l.(float64); ok {
		if rv, ok := r.(float64); ok {
			switch node.Operator {
			case ast.MULTI:
				return lv * rv, nil
			case ast.DIVIDE:
				return lv / rv, nil
			}
		}
		return nil, errors.New("right value is not number")
	}
	return nil, errors.New("left value is not number")
}

func evaluateAddition(node *ast.AdditionExpressionNode, env Env) (interface{}, error) {
	l, err := EvaluateExpression(node.Left, env)
	if err != nil {
		return nil, err
	}
	r, err := EvaluateExpression(node.Right, env)
	if err != nil {
		return nil, err
	}
	if lv, ok := l.(float64); ok {
		if rv, ok := r.(float64); ok {
			switch node.Operator {
			case ast.MINUS:
				return lv - rv, nil
			case ast.PLUS:
				return lv + rv, nil
			}
		}
		return nil, errors.New("right value is not number")
	} else if lv, ok := l.(string); ok {
		if rv, ok := r.(string); ok {
			switch node.Operator {
			case ast.MINUS:
				return nil, errors.New("string cannot substract")
			case ast.PLUS:
				return lv + rv, nil
			}
		}
		return nil, errors.New("right value is not string")
	}
	return nil, errors.New("Runtime Error. unexpected type")
}

func evaluateNotEqual(node *ast.NotEqualityExpressionNode, env Env) (interface{}, error) {
	l, err := EvaluateExpression(node.Left, env)
	if err != nil {
		return nil, err
	}
	r, err := EvaluateExpression(node.Right, env)
	if err != nil {
		return nil, err
	}
	b, err := evaluateEqualityInternal(l, r)
	return !b, err
}

func evaluateNot(node *ast.NotExpressionNode, env Env) (interface{}, error) {
	e, err := EvaluateExpression(node.Node, env)
	if err != nil {
		return nil, err
	}
	if v, ok := e.(bool); ok {
		return !v, nil
	}
	return nil, errors.New("Runtime Error. unexpected type")
}

func evaluateEqual(node *ast.EqualityExpressionNode, env Env) (interface{}, error) {
	l, err := EvaluateExpression(node.Left, env)
	if err != nil {
		return nil, err
	}
	r, err := EvaluateExpression(node.Right, env)
	if err != nil {
		return nil, err
	}
	return evaluateEqualityInternal(l, r)
}

func evaluateEqualityInternal(l interface{}, r interface{}) (bool, error) {
	if lv, ok := l.(float64); ok {
		if rv, ok := r.(float64); ok {
			return lv == rv, nil
		}
		return false, errors.New("right value is not number")
	} else if lv, ok := l.(string); ok {
		if rv, ok := r.(string); ok {
			return lv == rv, nil
		}
		return false, errors.New("right value is not string")
	} else if lv, ok := l.(bool); ok {
		if rv, ok := r.(bool); ok {
			return lv == rv, nil
		}
		return false, errors.New("right value is not bool")
	}
	return false, errors.New("Runtime Error. unexpected type")
}
