package parse

import (
	"bufio"

	"github.com/wreulicke/tincaml/ast"
)

func Parse(reader *bufio.Reader) (*ast.Tree, error) {
	l := &Lexer{}
	yyErrorVerbose = true
	l.Init(reader)
	yyParse(l)
	return l.result, l.error
}
