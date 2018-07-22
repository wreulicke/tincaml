package ast

import "fmt"

type FunctionCall struct {
	ID   ID
	Args []AST
}

func (n FunctionCall) String() string {
	return fmt.Sprintf("FunctionCall{ID: %s, Args: %v}", string(n.ID), n.Args)
}

type FunctionNode struct {
	ID     ID
	Params []Identifier
	Body   []AST
}

func (n FunctionNode) String() string {
	return fmt.Sprintf("FunctionNode(ID: %s, Params: %v, Body: %v)", string(n.ID), n.Params, n.Body)
}
