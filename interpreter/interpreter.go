package interpreter

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/wreulicke/tincaml/ast"
)

type Env map[string]interface{}

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
		fmt.Println(reflect.TypeOf(n), n)
		value, err := EvaluateExpression(n, env)
		if err != nil {
			return err
		}
		fmt.Println(value)
	}
	return nil
}

func EvaluateExpression(v ast.AST, env Env) (interface{}, error) {
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
	case *ast.FunctionCall:
		return evaluateFunctionCall(node, env)
	case *ast.FunctionNode:
		env[string(node.ID)] = node
		return node, nil
	default:
		return nil, errors.New("Unexpected condition. cannot evaluate this node")
	}
}

func evaluateFunctionCall(node *ast.FunctionCall, parentEnv Env) (interface{}, error) {
	v, ok := parentEnv[string(node.ID)]
	if !ok {
		return nil, fmt.Errorf("function: '%s' is not found", string(node.ID))
	}
	f, ok := v.(*ast.FunctionNode)
	if !ok {
		return nil, fmt.Errorf("'%s' is not function", string(node.ID))
	}
	if len(f.Params) != len(node.Args) {
		return nil, fmt.Errorf("size of parameters is expected to %d. but size of arguments is %d", len(f.Params), len(node.Args))
	}
	functionEnv := parentEnv.Clone()
	for k, v := range node.Args {
		id := f.Params[k]
		arg, err := EvaluateExpression(v, parentEnv)
		if err != nil {
			return nil, err
		}
		functionEnv[string(id.ID)] = arg
	}
	var val interface{}
	var err error
	for _, expr := range f.Body {
		val, err = EvaluateExpression(expr, functionEnv)
		if err != nil {
			return nil, err
		}
	}
	return val, err
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
