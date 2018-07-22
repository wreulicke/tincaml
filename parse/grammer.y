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

%nonassoc SEMICOLON
%right prec_let
%right prec_seq
%right prec_if
%right prec_fun
%left ASSIGN
%left EQUALITY NOT_EQUALITY
%left LESS GREATER LESS_EQUAL GREATER_EQUAL
%left MINUS PLUS
%left MULTI DIVIDE
%right UMINUS UNOT

%start start

%%

start: statements { 
    tree := &ast.Tree{$1}
    yylex.(*Lexer).result = tree
}

statements: 
    expression %prec prec_seq {
        $$ = []ast.AST{$1}
    } 
    | statements SEMICOLON expression  {
        values := make([]ast.AST, 0, len($1) + 1)
        values = append(values, $1...)
        values = append(values, $3)
        $$ = values
    }
     
expression: 
    '(' expression ')' {
        $$ = $2
    }
    | MINUS expression %prec UMINUS {
        $$ = &ast.NegativeNode{$2}
    }
    | NOT expression %prec UNOT {
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
    | 
    expression LESS expression { 
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
    | LET ID ASSIGN expression 
        %prec prec_let
        {
        $$ = &ast.AssignmentExpressionNode{
            ID: ast.ID($2.literal),
            Initializer: $4,
        }
    }
    | IF expression THEN statements ELSE statements 
        %prec prec_if {
        $$ = &ast.IfExpressionNode{
            Cond: $2,
            Then: $4,
            Else: $6,
        }
    }
    | primary_expression {
      $$ = $1
    }

fn_call: 
    expression '(' ')' {
        $$ = &ast.FunctionCall{
            Function: $1,
            Args: []ast.AST{},
        }  
    }
    | expression '(' arguments ')' {
        $$ = &ast.FunctionCall{
            Function: $1,
            Args: $3,
        }
    }

arguments: 
    expression ',' arguments {  
        values := make([]ast.AST, 0, len($3) + 1)
        values = append(values, $1)
        values = append(values, $3...)
        $$ = values 
    }
    | expression { 
        $$ = []ast.AST{ $1 } 
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
    ID params {
        values := make([]ast.Identifier, 0, len($2) + 1)
        values = append(values, ast.Identifier{ ast.ID($1.literal) })
        values = append(values, $2...)
        $$ = values
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