%{
package parse

import "github.com/wreulicke/tincaml/ast"

%}

%union{
    tree    *ast.Tree
    ast     ast.AST
    expr    ast.AST
    exprs    []ast.AST
    token   Token
}

%type<tree> start
%type<exprs> expressions
%type<expr> expression
%type<ast> number_literal primary_expression term
%token<token> NUMBER TRUE FALSE STRING
%token<token> MINUS PLUS MULTI DIVIDE ASSIGN EQUALITY NOT_EQUALITY NOT
%token<token> COLON

%left PLUS MINUS 
%left MULTI DIVIDE
%left EQUALITY ASSIGN
%right unary_minus
%right unary_not

%start start

%%

start: expressions { 
    tree := &ast.Tree{$1}
    yylex.(*Lexer).result = tree
    $$ = tree
}

expressions: 
    expression expressions {
        values := make([]ast.AST, 0, len($2) + 1)
        values = append(values, $1)
        values = append(values, $2...)
        $$ = values
    }
    | expression COLON expressions {
        values := make([]ast.AST, 0, len($3) + 1)
        values = append(values, $1)
        values = append(values, $3...)
        $$ = values
    }
    | expression {
        $$ = []ast.AST{$1}
    } 
     

expression: 
    expression PLUS expression {
        $$ = &ast.AdditionExpressionNode{
            Left: $1,
            Right: $3,
            Operator: ast.PLUS,
        }
    }
    | expression MINUS expression {
        $$ = &ast.AdditionExpressionNode{
            Left: $1,
            Right: $3,
            Operator: ast.MINUS,
        }
    }
    | expression MULTI expression {
        $$ = &ast.MultiplicativeExpressionNode{
            Left: $1,
            Right: $3,
            Operator: ast.MULTI,
        }
    }
    | expression DIVIDE expression {
        $$ = &ast.MultiplicativeExpressionNode{
            Left: $1,
            Right: $3,
            Operator: ast.DIVIDE,
        }
    }
    | expression EQUALITY expression {
        $$ = &ast.EqualityExpressionNode{
            Left: $1,
            Right: $3,
        }
    }
    | expression NOT_EQUALITY expression {
        $$ = &ast.NotEqualityExpressionNode{
            Left: $1,
            Right: $3,
        }
    }
    | MINUS expression %prec unary_minus {
        $$ = &ast.AdditionExpressionNode{
            Left: &ast.NumberNode{Value: 0},
            Right: $2,
            Operator: ast.MINUS,
        }
    }
    | NOT expression %prec unary_not {
        $$ = &ast.NotExpressionNode{$2}
    }
    | term {
      $$ = $1
    }
    
term: 
   primary_expression {
      $$ = $1
   }

primary_expression: 
    FALSE {
        $$ = &ast.BooleanNode{Value: false}
    }
    | TRUE {
        $$ = &ast.BooleanNode{Value: true}
    }
    | number_literal {
        $$ = $1
    }
    | STRING {
        $$ = &ast.StringNode{Value: $1.literal}
    }

number_literal: 
    NUMBER {
        lex := yylex.(*Lexer)
        num := lex.parseFloat($1.literal)
        $$ = &ast.NumberNode{Value: num}
    }

%%