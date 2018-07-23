%{
package parse

import "github.com/wreulicke/tincaml/ast"

%}

%union{
    tree    *ast.Tree
    ast     ast.AST
    expr    ast.AST
    exprs    []ast.AST
    params  []ast.Identifier
    token   Token
}

%type<> start
%type<exprs> statements arguments
%type<expr> expression fn_call fn_declare
%type<ast> primary_expression
%type<params> params
%token<token> NUMBER TRUE FALSE STRING ID
%token<token> MINUS PLUS MULTI DIVIDE ASSIGN EQUALITY NOT_EQUALITY NOT
%token<token> LESS GREATER LESS_EQUAL GREATER_EQUAL
%token<token> BEGIN_BLOCK END_BLOCK
%token<token> LET IF THEN ELSE
%token<token> SEMICOLON

%right prec_let
%right prec_seq
%right SEMICOLON
%right prec_if
%right prec_fun
%left ASSIGN EQUALITY NOT_EQUALITY
%left LESS GREATER LESS_EQUAL GREATER_EQUAL
%left PLUS MINUS
%left MULTI DIVIDE
%right prec_unary_minus
%left prec_app

%start start

%%

start: statements { 
    tree := &ast.Tree{$1}
    yylex.(*Lexer).result = tree
}

statements: 
    expression {
        $$ = []ast.AST{$1}
    } 
    | statements SEMICOLON expression  {
        $$ = append($1, $3)
    }
     
expression: 
    primary_expression {
      $$ = $1
    }
    | '(' expression ')' {
      $$ = $2
    }
    | MINUS expression %prec prec_unary_minus {
        $$ = &ast.NegativeNode{$2}
    }
    | NOT expression %prec prec_unary_minus {
        $$ = &ast.NotExpressionNode{$2}
    }
    | expression PLUS expression {
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
    | expression LESS expression { 
        $$ = &ast.RelationalExpressionNode{
            Left: $1, 
            Right: $3,
            Operator: ast.LESS,
        } 
    }
    | expression GREATER expression { 
        $$ = &ast.RelationalExpressionNode{
            Left: $1, 
            Right: $3,
            Operator: ast.GREATER,
        } 
    }
    | expression LESS_EQUAL expression { 
        $$ = &ast.RelationalExpressionNode{
            Left: $1, 
            Right: $3,
            Operator: ast.LESS_EQUAL,
        } 
    }
    | expression GREATER_EQUAL expression { 
        $$ = &ast.RelationalExpressionNode{
            Left: $1, 
            Right: $3,
            Operator: ast.GREATER_EQUAL,
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
    | fn_call {
        $$ = $1
    }
    | fn_declare {
        $$ = $1
    }
    | LET ID ASSIGN expression %prec prec_let {
        $$ = &ast.AssignmentExpressionNode{
            ID: ast.ID($2.literal),
            Initializer: $4,
        }
    }
    | IF expression THEN statements ELSE statements %prec prec_if {
        $$ = &ast.IfExpressionNode{
            Cond: $2,
            Then: $4,
            Else: $6,
        }
    }

fn_call: 
    expression '(' ')' %prec prec_fun {
        $$ = &ast.FunctionCall{
            Function: $1,
            Args: []ast.AST{},
        }  
    }
    | expression '(' expression ')' %prec prec_fun {
        $$ = &ast.FunctionCall{
            Function: $1,
            Args: []ast.AST{$3},
        }
    }
    | expression '(' arguments ')' %prec prec_fun {
        $$ = &ast.FunctionCall{
            Function: $1,
            Args: $3,
        }
    }

arguments: 
    arguments ',' expression {  
        $$ = append($1, $3) 
    }
    | expression ',' expression {
        $$ = []ast.AST{$1, $3} 
    }

fn_declare: 
    LET ID params BEGIN_BLOCK statements END_BLOCK {
        $$ = &ast.FunctionNode{
            ID: ast.ID($2.literal),
            Params: $3,
            Body: $5,
        }
    }
    | LET ID BEGIN_BLOCK statements END_BLOCK {
        $$ = &ast.FunctionNode{
            ID: ast.ID($2.literal),
            Params: []ast.Identifier{},
            Body: $4,
        }
    }

params: 
    params ID {
        $$ = append($1, ast.Identifier{ ast.ID($2.literal) })
    }
    | ID { 
        $$ = []ast.Identifier{ ast.Identifier{ ast.ID($1.literal) } } 
    }

primary_expression:
    FALSE {
        $$ = &ast.BooleanNode{Value: false}
    }
    | TRUE {
        $$ = &ast.BooleanNode{Value: true}
    }
    | NUMBER {
        lex := yylex.(*Lexer)
        num := lex.parseFloat($1.literal)
        $$ = &ast.NumberNode{Value: num}
    }
    | STRING {
        $$ = &ast.StringNode{$1.literal}
    }
    | ID {
        $$ = &ast.Identifier{ ast.ID($1.literal) }
    }
%%