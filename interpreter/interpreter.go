package interpreter

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/wreulicke/tincaml/ast"
)

func Evaluate(tree *ast.Tree) error {
	for _, n := range tree.Body {
		fmt.Println(reflect.TypeOf(n), n)
		value, err := EvaluateExpression(n)
		if err != nil {
			return err
		}
		fmt.Println(value)
	}
	return nil
}

func EvaluateExpression(v ast.AST) (interface{}, error) {
	switch node := v.(type) {
	case *ast.AdditionExpressionNode:
		return evaluateAddition(node)
	case *ast.MultiplicativeExpressionNode:
		return evaluateMultiplicative(node)
	case *ast.EqualityExpressionNode:
		return evaluateEqual(node)
	case *ast.NotEqualityExpressionNode:
		return evaluateNotEqual(node)
	case *ast.NotExpressionNode:
		return evaluateNot(node)
	case *ast.BooleanNode:
		return node.Value, nil
	case *ast.NumberNode:
		return node.Value, nil
	case *ast.StringNode:
		return node.Value, nil
	default:
		return nil, errors.New("Unexpected condition. cannot evaluate this node")
	}
}

func evaluateMultiplicative(node *ast.MultiplicativeExpressionNode) (interface{}, error) {
	l, err := EvaluateExpression(node.Left)
	if err != nil {
		return nil, err
	}
	r, err := EvaluateExpression(node.Right)
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

func evaluateAddition(node *ast.AdditionExpressionNode) (interface{}, error) {
	l, err := EvaluateExpression(node.Left)
	if err != nil {
		return nil, err
	}
	r, err := EvaluateExpression(node.Right)
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

func evaluateNotEqual(node *ast.NotEqualityExpressionNode) (interface{}, error) {
	l, err := EvaluateExpression(node.Left)
	if err != nil {
		return nil, err
	}
	r, err := EvaluateExpression(node.Right)
	if err != nil {
		return nil, err
	}
	b, err := evaluateEqualityInternal(l, r)
	return !b, err
}

func evaluateNot(node *ast.NotExpressionNode) (interface{}, error) {
	e, err := EvaluateExpression(node.Node)
	if err != nil {
		return nil, err
	}
	if v, ok := e.(bool); ok {
		return !v, nil
	}
	return nil, errors.New("Runtime Error. unexpected type")
}

func evaluateEqual(node *ast.EqualityExpressionNode) (interface{}, error) {
	l, err := EvaluateExpression(node.Left)
	if err != nil {
		return nil, err
	}
	r, err := EvaluateExpression(node.Right)
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
