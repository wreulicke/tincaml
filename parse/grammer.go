// Code generated by goyacc -o grammer.go grammer.y. DO NOT EDIT.

//line grammer.y:2
package parse

import __yyfmt__ "fmt"

//line grammer.y:2
import "github.com/wreulicke/tincaml/ast"

//line grammer.y:8
type yySymType struct {
	yys   int
	tree  *ast.Tree
	ast   ast.AST
	expr  ast.AST
	exprs []ast.AST
	token Token
}

const NUMBER = 57346
const TRUE = 57347
const FALSE = 57348
const STRING = 57349
const MINUS = 57350
const PLUS = 57351
const MULTI = 57352
const DIVIDE = 57353
const ASSIGN = 57354
const EQUALITY = 57355
const NOT_EQUALITY = 57356
const NOT = 57357
const COLON = 57358
const UMINUS = 57359
const UNOT = 57360

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"NUMBER",
	"TRUE",
	"FALSE",
	"STRING",
	"MINUS",
	"PLUS",
	"MULTI",
	"DIVIDE",
	"ASSIGN",
	"EQUALITY",
	"NOT_EQUALITY",
	"NOT",
	"COLON",
	"UMINUS",
	"UNOT",
	"'('",
	"')'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line grammer.y:134

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 45

var yyAct = [...]int{

	4, 14, 16, 15, 17, 18, 7, 21, 5, 22,
	23, 2, 17, 18, 3, 1, 25, 26, 27, 28,
	29, 30, 16, 15, 17, 18, 24, 19, 20, 12,
	11, 10, 13, 8, 31, 0, 16, 15, 17, 18,
	9, 19, 20, 0, 6,
}
var yyPact = [...]int{

	25, -1000, -1000, -15, 28, -1000, 25, -1000, 25, 25,
	-1000, -1000, -1000, -1000, 25, 25, 25, 25, 25, 25,
	25, 14, -1000, -1000, -1000, 2, 2, -1000, -1000, -6,
	-6, -1000,
}
var yyPgo = [...]int{

	0, 15, 11, 14, 0, 8, 6,
}
var yyR1 = [...]int{

	0, 1, 2, 2, 3, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 5, 5, 6, 6, 6, 6,
}
var yyR2 = [...]int{

	0, 1, 3, 1, 1, 1, 3, 3, 3, 3,
	3, 3, 3, 1, 2, 2, 1, 1, 1, 1,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, -4, -5, 19, -6, 8, 15,
	6, 5, 4, 7, 16, 9, 8, 10, 11, 13,
	14, -4, -4, -4, -2, -4, -4, -4, -4, -4,
	-4, 20,
}
var yyDef = [...]int{

	0, -2, 1, 3, 4, 5, 0, 13, 0, 0,
	16, 17, 18, 19, 0, 0, 0, 0, 0, 0,
	0, 0, 14, 15, 2, 7, 8, 9, 10, 11,
	12, 6,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	19, 20,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18,
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
		//line grammer.y:34
		{
			tree := &ast.Tree{yyDollar[1].exprs}
			yylex.(*Lexer).result = tree
		}
	case 2:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammer.y:40
		{
			values := make([]ast.AST, 0, len(yyDollar[3].exprs)+1)
			values = append(values, yyDollar[1].expr)
			values = append(values, yyDollar[3].exprs...)
			yyVAL.exprs = values
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammer.y:46
		{
			yyVAL.exprs = []ast.AST{yyDollar[1].expr}
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammer.y:52
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammer.y:57
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 6:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammer.y:60
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 7:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammer.y:63
		{
			yyVAL.expr = &ast.AdditionExpressionNode{
				Left:     yyDollar[1].expr,
				Right:    yyDollar[3].expr,
				Operator: ast.PLUS,
			}
		}
	case 8:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammer.y:70
		{
			yyVAL.expr = &ast.AdditionExpressionNode{
				Left:     yyDollar[1].expr,
				Right:    yyDollar[3].expr,
				Operator: ast.MINUS,
			}
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammer.y:77
		{
			yyVAL.expr = &ast.MultiplicativeExpressionNode{
				Left:     yyDollar[1].expr,
				Right:    yyDollar[3].expr,
				Operator: ast.MULTI,
			}
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammer.y:84
		{
			yyVAL.expr = &ast.MultiplicativeExpressionNode{
				Left:     yyDollar[1].expr,
				Right:    yyDollar[3].expr,
				Operator: ast.DIVIDE,
			}
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammer.y:91
		{
			yyVAL.expr = &ast.EqualityExpressionNode{
				Left:  yyDollar[1].expr,
				Right: yyDollar[3].expr,
			}
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammer.y:97
		{
			yyVAL.expr = &ast.NotEqualityExpressionNode{
				Left:  yyDollar[1].expr,
				Right: yyDollar[3].expr,
			}
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammer.y:103
		{
			yyVAL.expr = yyDollar[1].ast
		}
	case 14:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line grammer.y:108
		{
			yyVAL.expr = &ast.AdditionExpressionNode{
				Left:     &ast.NumberNode{Value: 0},
				Right:    yyDollar[2].expr,
				Operator: ast.MINUS,
			}
		}
	case 15:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line grammer.y:115
		{
			yyVAL.expr = &ast.NotExpressionNode{yyDollar[2].expr}
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammer.y:120
		{
			yyVAL.ast = &ast.BooleanNode{Value: false}
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammer.y:123
		{
			yyVAL.ast = &ast.BooleanNode{Value: true}
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammer.y:126
		{
			lex := yylex.(*Lexer)
			num := lex.parseFloat(yyDollar[1].token.literal)
			yyVAL.ast = &ast.NumberNode{Value: num}
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammer.y:131
		{
			yyVAL.ast = &ast.StringNode{Value: yyDollar[1].token.literal}
		}
	}
	goto yystack /* stack new state and value */
}
