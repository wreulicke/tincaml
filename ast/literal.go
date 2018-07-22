package ast

import (
	"fmt"
)

type NumberNode struct {
	Value float64
}

func (n NumberNode) String() string {
	return fmt.Sprintf("NumberNode(%f)", n.Value)
}

type StringNode struct {
	Value string
}

func (n StringNode) String() string {
	return fmt.Sprintf("StringNode(%s)", n.Value)
}

type BooleanNode struct {
	Value bool
}

func (n BooleanNode) String() string {
	return fmt.Sprintf("BooleanNode(%v)", n.Value)
}

type ID string

type Identifier struct {
	ID ID
}

func (n Identifier) String() string {
	return fmt.Sprintf("Identifier(%s)", string(n.ID))
}
