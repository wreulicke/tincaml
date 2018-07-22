package ast

import "fmt"

type FunctionCall struct {
	Function AST
	Args     []AST
}

func (n FunctionCall) String() string {
	return fmt.Sprintf("FunctionCall{Function: %s, Args: %v}", n.Function, n.Args)
}

type FunctionNode struct {
	ID     ID
	Params []Identifier
	Body   []AST
}

func (n FunctionNode) String() string {
	return fmt.Sprintf("FunctionNode(ID: %s, Params: %v, Body: %v)", string(n.ID), n.Params, n.Body)
}
