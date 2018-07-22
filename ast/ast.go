package ast

import (
	"fmt"
)

type ArrayNode struct {
	Value []AST
}
type NegativeNode struct {
	Node AST
}

func (n NegativeNode) String() string {
	return fmt.Sprintf("NegativeNode(%v)", n.Node)
}

type EmptyExpressionNode struct {
}

func (n EmptyExpressionNode) String() string {
	return "EmptyExpressionNode"
}

type AdditionOperator int

const (
	// PLUS is '+'
	PLUS = iota
	// MINUS is '-'
	MINUS
)

type AdditionExpressionNode struct {
	Left     AST
	Right    AST
	Operator AdditionOperator
}

func (n AdditionExpressionNode) String() string {
	switch n.Operator {
	case PLUS:
		return fmt.Sprintf("AdditionExpressionNode(%v + %v)", n.Left, n.Right)
	case MINUS:
		return fmt.Sprintf("AdditionExpressionNode(%v - %v)", n.Left, n.Right)
	}
	panic("cannot reach here")
}

type MultiplicativeOperator int

const (
	// MULTI is '*'
	MULTI = iota
	// DIVIDE is '/'
	DIVIDE
)

type MultiplicativeExpressionNode struct {
	Left     AST
	Right    AST
	Operator MultiplicativeOperator
}

func (n MultiplicativeExpressionNode) String() string {
	switch n.Operator {
	case MULTI:
		return fmt.Sprintf("MultiplicativeExpressionNode(%v * %v)", n.Left, n.Right)
	case DIVIDE:
		return fmt.Sprintf("MultiplicativeExpressionNode(%v / %v)", n.Left, n.Right)
	}
	panic("cannot reach here")
}

type EqualityExpressionNode struct {
	Left  AST
	Right AST
}

func (n EqualityExpressionNode) String() string {
	return fmt.Sprintf("EqualityExpressionNode(%v == %v)", n.Left, n.Right)
}

type NotEqualityExpressionNode struct {
	Left  AST
	Right AST
}

func (n NotEqualityExpressionNode) String() string {
	return fmt.Sprintf("NotEqualityExpressionNode(%v != %v)", n.Left, n.Right)
}

type NotExpressionNode struct {
	Node AST
}

func (n NotExpressionNode) String() string {
	return fmt.Sprintf("NotExpressionNode(%v)", n.Node)
}

type IfExpressionNode struct {
	Cond AST
	Then []AST
	Else []AST
}

func (n IfExpressionNode) String() string {
	return fmt.Sprintf("IfExpressionNode(Cond: %v, Then: %v, Else: %v)", n.Cond, n.Then, n.Else)
}

type RelationalOperator int

const (
	LESS = iota
	GREATER
	LESS_EQUAL
	GREATER_EQUAL
)

type RelationalExpressionNode struct {
	Left     AST
	Right    AST
	Operator RelationalOperator
}

func (n RelationalExpressionNode) String() string {
	switch n.Operator {
	case LESS:
		return fmt.Sprintf("RelationalExpressionNode(%v < %v)", n.Left, n.Right)
	case GREATER:
		return fmt.Sprintf("RelationalExpressionNode(%v > %v)", n.Left, n.Right)
	case LESS_EQUAL:
		return fmt.Sprintf("RelationalExpressionNode(%v <= %v)", n.Left, n.Right)
	case GREATER_EQUAL:
		return fmt.Sprintf("RelationalExpressionNode(%v >= %v)", n.Left, n.Right)
	}
	panic("cannot reach here")
}

type AssignmentExpressionNode struct {
	ID          ID
	Initializer AST
}

func (n AssignmentExpressionNode) String() string {
	return fmt.Sprintf("AssignmentExpressionNode(%s = %v)", string(n.ID), n.Initializer)
}

type AST interface {
	String() string
}

type Tree struct {
	Body []AST
}
