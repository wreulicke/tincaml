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
%type<expr> statement expression unary_expression fn_call fn_declare if_expr relational_expr
%type<ast> primary_expression
%type<params> params
%token<token> NUMBER TRUE FALSE STRING ID
%token<token> MINUS PLUS MULTI DIVIDE ASSIGN EQUALITY NOT_EQUALITY NOT
%token<token> LESS GREATER LESS_EQUAL GREATER_EQUAL
%token<token> BEGIN_BLOCK END_BLOCK
%token<token> LET IF THEN ELSE
%token<token> COLON

%nonassoc COLON
%right prec_if
%left EQUALITY NOT_EQUALITY ASSIGN
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
    statement COLON statements {
        values := make([]ast.AST, 0, len($3) + 1)
        values = append(values, $1)
        values = append(values, $3...)
        $$ = values
    }
    | statement COLON {
        $$ = []ast.AST{$1}
    } 
    | statement {
        $$ = []ast.AST{$1}
    } 
     

statement: 
    expression {
        $$ = $1
    }
    | if_expr %prec prec_if {
        $$ = $1
    }

expression: 
    unary_expression %prec UMINUS {
        $$ = $1
    }
    | '(' expression ')' {
        $$ = $2
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
    | relational_expr {
        $$ = $1
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
    | primary_expression {
      $$ = $1
    }

relational_expr:
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

unary_expression:
    MINUS expression %prec UMINUS {
        $$ = &ast.AdditionExpressionNode{
            Left: &ast.NumberNode{Value: 0},
            Right: $2,
            Operator: ast.MINUS,
        }
    }
    | NOT expression %prec UNOT {
        $$ = &ast.NotExpressionNode{$2}
    }

if_expr:
    IF expression THEN statements ELSE statements {
        $$ = &ast.IfExpressionNode{
            Cond: $2,
            Then: $4,
            Else: $6,
        }
    }

fn_call: 
    ID '(' arguments ')' {
        $$ = &ast.FunctionCall{
            ID: ast.ID($1.literal),
            Args: $3,
        }   
    }
    | ID '(' ')' {
        $$ = &ast.FunctionCall{
            ID: ast.ID($1.literal),
            Args: []ast.AST{},
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