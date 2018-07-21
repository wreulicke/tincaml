// Code generated by goyacc -o grammer.go grammer.y. DO NOT EDIT.

//line grammer.y:2
package parse

import __yyfmt__ "fmt"

//line grammer.y:2
import "github.com/wreulicke/tincaml/ast"

//line grammer.y:8
type yySymType struct {
	yys    int
	tree   *ast.Tree
	ast    ast.AST
	expr   ast.AST
	exprs  []ast.AST
	params []ast.Identifier
	token  Token
}

const NUMBER = 57346
const TRUE = 57347
const FALSE = 57348
const STRING = 57349
const ID = 57350
const MINUS = 57351
const PLUS = 57352
const MULTI = 57353
const DIVIDE = 57354
const ASSIGN = 57355
const EQUALITY = 57356
const NOT_EQUALITY = 57357
const NOT = 57358
const LESS = 57359
const GREATER = 57360
const LESS_EQUAL = 57361
const GREATER_EQUAL = 57362
const BEGIN_BLOCK = 57363
const END_BLOCK = 57364
const LET = 57365
const IF = 57366
const THEN = 57367
const ELSE = 57368
const COLON = 57369
const prec_if = 57370
const UMINUS = 57371
const UNOT = 57372

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"NUMBER",
	"TRUE",
	"FALSE",
	"STRING",
	"ID",
	"MINUS",
	"PLUS",
	"MULTI",
	"DIVIDE",
	"ASSIGN",
	"EQUALITY",
	"NOT_EQUALITY",
	"NOT",
	"LESS",
	"GREATER",
	"LESS_EQUAL",
	"GREATER_EQUAL",
	"BEGIN_BLOCK",
	"END_BLOCK",
	"LET",
	"IF",
	"THEN",
	"ELSE",
	"COLON",
	"prec_if",
	"UMINUS",
	"UNOT",
	"'('",
	"')'",
	"','",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line grammer.y:250

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 170

var yyAct = [...]int{

	4, 51, 54, 58, 36, 21, 63, 68, 32, 66,
	60, 2, 56, 33, 34, 35, 23, 22, 24, 25,
	37, 56, 11, 39, 40, 41, 42, 43, 44, 45,
	46, 47, 48, 38, 55, 24, 25, 53, 8, 23,
	22, 24, 25, 5, 26, 27, 10, 28, 29, 30,
	31, 9, 6, 3, 1, 0, 0, 0, 0, 62,
	53, 64, 57, 59, 0, 0, 0, 61, 23, 22,
	24, 25, 65, 26, 27, 67, 28, 29, 30, 31,
	19, 18, 17, 20, 15, 13, 0, 0, 0, 0,
	0, 49, 14, 19, 18, 17, 20, 15, 13, 16,
	0, 0, 0, 0, 0, 14, 0, 7, 52, 0,
	0, 0, 16, 12, 19, 18, 17, 20, 15, 13,
	7, 0, 0, 0, 0, 0, 14, 0, 0, 23,
	22, 24, 25, 16, 26, 27, 0, 28, 29, 30,
	31, 7, 0, 0, 0, 50, 23, 22, 24, 25,
	0, 26, 27, 0, 28, 29, 30, 31, 23, 22,
	24, 25, 0, 0, 0, 0, 28, 29, 30, 31,
}
var yyPact = [...]int{

	89, -1000, -1000, -22, 137, -1000, -1000, 110, -1000, -1000,
	-1000, -1000, 110, 110, 110, -27, 12, -1000, -1000, -1000,
	-1000, 89, 110, 110, 110, 110, 110, 110, 110, 110,
	110, 110, 59, 120, -1000, -1000, 76, 13, -1000, 24,
	24, -1000, -1000, 149, 149, 7, 7, 7, 7, -1000,
	89, -29, -1000, 30, -11, 89, 4, -20, -1000, 110,
	89, -13, -1000, 89, -1000, -15, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 54, 11, 1, 53, 0, 52, 51, 46, 43,
	38, 22, 2,
}
var yyR1 = [...]int{

	0, 1, 2, 2, 2, 4, 4, 5, 5, 5,
	5, 5, 5, 5, 5, 5, 5, 5, 5, 10,
	10, 10, 10, 6, 6, 9, 7, 7, 3, 3,
	8, 8, 12, 12, 11, 11, 11, 11, 11,
}
var yyR2 = [...]int{

	0, 1, 3, 2, 1, 1, 1, 1, 3, 3,
	3, 3, 1, 3, 3, 3, 1, 1, 1, 3,
	3, 3, 3, 2, 2, 6, 4, 3, 3, 1,
	6, 5, 2, 1, 1, 1, 1, 1, 1,
}
var yyChk = [...]int{

	-1000, -1, -2, -4, -5, -9, -6, 31, -10, -7,
	-8, -11, 24, 9, 16, 8, 23, 6, 5, 4,
	7, 27, 10, 9, 11, 12, 14, 15, 17, 18,
	19, 20, -5, -5, -5, -5, 31, 8, -2, -5,
	-5, -5, -5, -5, -5, -5, -5, -5, -5, 32,
	25, -3, 32, -5, -12, 21, 8, -2, 32, 33,
	21, -2, -12, 26, -3, -2, 22, -2, 22,
}
var yyDef = [...]int{

	0, -2, 1, 4, 5, 6, 7, 0, 12, 16,
	17, 18, 0, 0, 0, 38, 0, 34, 35, 36,
	37, 3, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 23, 24, 0, 0, 2, 9,
	10, 11, 13, 14, 15, 19, 20, 21, 22, 8,
	0, 0, 27, 29, 0, 0, 33, 0, 26, 0,
	0, 0, 32, 0, 28, 0, 31, 25, 30,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	31, 32, 3, 3, 33,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammer.y:41
		{
			tree := &ast.Tree{yyDollar[1].exprs}
			yylex.(*Lexer).result = tree
		}
	case 2:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammer.y:47
		{
			values := make([]ast.AST, 0, len(yyDollar[3].exprs)+1)
			values = append(values, yyDollar[1].expr)
			values = append(values, yyDollar[3].exprs...)
			yyVAL.exprs = values
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line grammer.y:53
		{
			yyVAL.exprs = []ast.AST{yyDollar[1].expr}
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammer.y:56
		{
			yyVAL.exprs = []ast.AST{yyDollar[1].expr}
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammer.y:62
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammer.y:65
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammer.y:70
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 8:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammer.y:73
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammer.y:76
		{
			yyVAL.expr = &ast.AdditionExpressionNode{
				Left:     yyDollar[1].expr,
				Right:    yyDollar[3].expr,
				Operator: ast.PLUS,
			}
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammer.y:83
		{
			yyVAL.expr = &ast.AdditionExpressionNode{
				Left:     yyDollar[1].expr,
				Right:    yyDollar[3].expr,
				Operator: ast.MINUS,
			}
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammer.y:90
		{
			yyVAL.expr = &ast.MultiplicativeExpressionNode{
				Left:     yyDollar[1].expr,
				Right:    yyDollar[3].expr,
				Operator: ast.MULTI,
			}
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammer.y:97
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammer.y:100
		{
			yyVAL.expr = &ast.MultiplicativeExpressionNode{
				Left:     yyDollar[1].expr,
				Right:    yyDollar[3].expr,
				Operator: ast.DIVIDE,
			}
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammer.y:107
		{
			yyVAL.expr = &ast.EqualityExpressionNode{
				Left:  yyDollar[1].expr,
				Right: yyDollar[3].expr,
			}
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammer.y:113
		{
			yyVAL.expr = &ast.NotEqualityExpressionNode{
				Left:  yyDollar[1].expr,
				Right: yyDollar[3].expr,
			}
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammer.y:119
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammer.y:122
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammer.y:125
		{
			yyVAL.expr = yyDollar[1].ast
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammer.y:130
		{
			yyVAL.expr = &ast.RelationalExpressionNode{
				Left:     yyDollar[1].expr,
				Right:    yyDollar[3].expr,
				Operator: ast.LESS,
			}
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammer.y:137
		{
			yyVAL.expr = &ast.RelationalExpressionNode{
				Left:     yyDollar[1].expr,
				Right:    yyDollar[3].expr,
				Operator: ast.GREATER,
			}
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammer.y:144
		{
			yyVAL.expr = &ast.RelationalExpressionNode{
				Left:     yyDollar[1].expr,
				Right:    yyDollar[3].expr,
				Operator: ast.LESS_EQUAL,
			}
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammer.y:151
		{
			yyVAL.expr = &ast.RelationalExpressionNode{
				Left:     yyDollar[1].expr,
				Right:    yyDollar[3].expr,
				Operator: ast.GREATER_EQUAL,
			}
		}
	case 23:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line grammer.y:160
		{
			yyVAL.expr = &ast.AdditionExpressionNode{
				Left:     &ast.NumberNode{Value: 0},
				Right:    yyDollar[2].expr,
				Operator: ast.MINUS,
			}
		}
	case 24:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line grammer.y:167
		{
			yyVAL.expr = &ast.NotExpressionNode{yyDollar[2].expr}
		}
	case 25:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line grammer.y:172
		{
			yyVAL.expr = &ast.IfExpressionNode{
				Cond: yyDollar[2].expr,
				Then: yyDollar[4].exprs,
				Else: yyDollar[6].exprs,
			}
		}
	case 26:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line grammer.y:181
		{
			yyVAL.expr = &ast.FunctionCall{
				ID:   ast.ID(yyDollar[1].token.literal),
				Args: yyDollar[3].exprs,
			}
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammer.y:187
		{
			yyVAL.expr = &ast.FunctionCall{
				ID:   ast.ID(yyDollar[1].token.literal),
				Args: []ast.AST{},
			}
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammer.y:195
		{
			values := make([]ast.AST, 0, len(yyDollar[3].exprs)+1)
			values = append(values, yyDollar[1].expr)
			values = append(values, yyDollar[3].exprs...)
			yyVAL.exprs = values
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammer.y:201
		{
			yyVAL.exprs = []ast.AST{yyDollar[1].expr}
		}
	case 30:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line grammer.y:206
		{
			yyVAL.expr = &ast.FunctionNode{
				ID:     ast.ID(yyDollar[2].token.literal),
				Params: yyDollar[3].params,
				Body:   yyDollar[5].exprs,
			}
		}
	case 31:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line grammer.y:213
		{
			yyVAL.expr = &ast.FunctionNode{
				ID:     ast.ID(yyDollar[2].token.literal),
				Params: []ast.Identifier{},
				Body:   yyDollar[4].exprs,
			}
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line grammer.y:222
		{
			values := make([]ast.Identifier, 0, len(yyDollar[2].params)+1)
			values = append(values, ast.Identifier{ast.ID(yyDollar[1].token.literal)})
			values = append(values, yyDollar[2].params...)
			yyVAL.params = values
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammer.y:228
		{
			yyVAL.params = []ast.Identifier{ast.Identifier{ast.ID(yyDollar[1].token.literal)}}
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammer.y:233
		{
			yyVAL.ast = &ast.BooleanNode{Value: false}
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammer.y:236
		{
			yyVAL.ast = &ast.BooleanNode{Value: true}
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammer.y:239
		{
			lex := yylex.(*Lexer)
			num := lex.parseFloat(yyDollar[1].token.literal)
			yyVAL.ast = &ast.NumberNode{Value: num}
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammer.y:244
		{
			yyVAL.ast = &ast.StringNode{yyDollar[1].token.literal}
		}
	case 38:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammer.y:247
		{
			yyVAL.ast = &ast.Identifier{ast.ID(yyDollar[1].token.literal)}
		}
	}
	goto yystack /* stack new state and value */
}
