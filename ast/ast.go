package ast

type NumberNode struct {
	Value float64
}

type StringNode struct {
	Value string
}

type BooleanNode struct {
	Value bool
}

type ID string

type Identifier struct {
	ID ID
}

type ArrayNode struct {
	Value []AST
}

type FunctionCall struct {
	ID   ID
	Args []AST
}

type FunctionNode struct {
	ID     ID
	Params []Identifier
	Body   []AST
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

type EqualityExpressionNode struct {
	Left  AST
	Right AST
}

type NotEqualityExpressionNode struct {
	Left  AST
	Right AST
}

type NotExpressionNode struct {
	Node AST
}

type IfExpressionNode struct {
	Cond AST
	Then []AST
	Else []AST
}

type AST interface {
}

type Tree struct {
	Body []AST
}
