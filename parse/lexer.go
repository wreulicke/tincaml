package parse

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf16"
	"unicode/utf8"

	"github.com/wreulicke/tincaml/ast"
)

type Token struct {
	typ     int
	literal string
}

func (token Token) String() string {
	return string(token.typ) + ":" + token.literal
}

type Position struct {
	line   int
	column int
}

type Lexer struct {
	input    *bufio.Reader
	buffer   bytes.Buffer
	position *Position
	offset   int
	result   *ast.Tree
	error    error
}

const eof = -1

func (l *Lexer) Init(reader io.Reader) {
	l.input = bufio.NewReader(reader)
	l.position = &Position{line: 1}
}

//go:generate goyacc -o grammer.go grammer.y
func (l *Lexer) Error(e string) {
	err := fmt.Errorf("%s in %d:%d", e, (*l).position.line, (*l).position.column)
	l.error = err
}

func (l *Lexer) parseFloat(str string) float64 {
	f64, err := strconv.ParseFloat(str, 64)
	if err != nil {
		l.Error("unexpected number format error")
		return -1
	}
	return f64
}

func (l *Lexer) scanDigit(next rune) {
	if next == '0' {
		l.Error("unexpected digit '0'")
		return
	} else if isDigit(next) {
		next := l.Peek()
		for {
			if !isDigit(next) {
				break
			}
			l.Next()
			next = l.Peek()
		}
		next = l.Peek()
		if next == '.' {
			l.Next()
			next = l.Peek()
			if !isDigit(next) {
				l.Error("unexpected token: expected digits")
				return
			}
			for {
				if !isDigit(next) {
					break
				}
				l.Next()
				next = l.Peek()
			}
		}
		next = l.Peek()
		if next == 'e' || next == 'E' {
			l.Next()
			next := l.Peek()
			if next == '+' || next == '-' {
				l.Next()
			}
			next = l.Peek()
			if !isDigit(next) {
				l.Error("digit expected for number exponent")
				return
			}
			l.Next()
			next = l.Peek()
			for {
				if !isDigit(next) {
					break
				}
				l.Next()
				next = l.Peek()
			}
		}
	} else {
		l.Error("error")
		return
	}
}

func isIdentifierPart(r rune) bool {
	return (unicode.IsLetter(r) || unicode.IsMark(r) || unicode.IsDigit(r) ||
		unicode.IsPunct(r)) && !strings.ContainsRune("{}[]():,", r)
}

func (l *Lexer) scanIdentifier() {
	next := l.Peek()
	if unicode.IsLetter(next) || next == '$' || next == '_' {
	} else {
		return
	}

	for next = l.Peek(); isIdentifierPart(next); {
		l.Next()
		next = l.Peek()
	}
}

func (l *Lexer) scanMultilineString() {
	for {
		switch next := l.Peek(); {
		case next == '`':
			l.Skip()
			return
		case next == eof:
			l.Error("unclosed string")
			return
		default:
			l.Next()
		}
	}
}

func (l *Lexer) scanString(start rune) {
	for {
		next := l.Peek()
		if next == start {
			l.Skip()
			return
		}
		switch {
		case next == '\\':
			l.Skip()
			next := l.Peek()
			if next == start {
				l.Next()
			} else if next == 'b' {
				l.Skip()
				l.buffer.WriteRune('\b')
			} else if next == 'f' {
				l.Skip()
				l.buffer.WriteRune('\f')
			} else if next == 'n' {
				l.Skip()
				l.buffer.WriteRune('\n')
			} else if next == 'r' {
				l.Skip()
				l.buffer.WriteRune('\r')
			} else if next == 't' {
				l.Skip()
				l.buffer.WriteRune('\t')
			} else if r := l.Peek(); r == 'u' {
				l.Skip()
				bytes := ""
				for i := 0; i < 4; i++ {
					b := l.Peek()
					if strings.IndexRune("0123456789ABDEFabcdef", l.Peek()) >= 0 {
						bytes = bytes + string(b)
						l.Skip()
					} else {
						l.Error("expected 4 hexadecimal digits")
						return
					}
				}
				b, _ := strconv.ParseUint(bytes, 16 /* hex */, 16 /* 2 bytes */)
				if utf16.IsSurrogate(rune(b)) {
					if l.Peek() == '\\' {
						l.Skip()
						if l.Peek() == 'u' {
							l.Skip()
							bytes := ""
							for i := 0; i < 4; i++ {
								b := l.Peek()
								if strings.IndexRune("0123456789ABDEFabcdef", l.Peek()) >= 0 {
									bytes = bytes + string(b)
									l.Skip()
								} else {
									l.Error("expected 4 hexadecimal digits")
									return
								}
							}
							b2, _ := strconv.ParseUint(bytes, 16 /* hex */, 16 /* 2 bytes */)
							l.buffer.WriteRune(utf16.DecodeRune(rune(b), rune(b2)))
						}
					} else {
						l.Error("invalid surrogate pair")
						return
					}
				} else {
					l.buffer.WriteRune(rune(b))
				}
			} else {
				l.Error("unsupported escape character")
				return
			}
		case unicode.IsControl(next):
			l.Error("cannot contain control characters in strings")
			return
		case next == eof:
			l.Error("unclosed string")
			return
		default:
			l.Next()
		}
	}
}

func (l *Lexer) scanWhitespace() {
	ruNe := l.Peek()
	for unicode.IsSpace(ruNe) {
		l.Next()
		ruNe = l.Peek()
	}
	return
}

func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) TokenText() string {
	return l.buffer.String()
}

func (l *Lexer) Scan() int {
retry:
	next := l.Peek()
	if next == '`' {
		l.Skip()
		l.scanMultilineString()
		return STRING
	} else if next == '\'' {
		l.Skip()
		l.scanString('\'')
		return STRING
	} else if next == '"' {
		l.Skip()
		l.scanString('"')
		return STRING
	}
	l.Next()
	switch {
	case next == '=':
		if l.Peek() == '=' {
			l.Next()
			return EQUALITY
		}
		return ASSIGN
	case next == '!':
		if l.Peek() == '=' {
			l.Next()
			return NOT_EQUALITY
		}
		return NOT
	case next == ';':
		l.scanWhitespace()
		return COLON
	case next == '-':
		return MINUS
	case next == ',':
		return int(',')
	case next == '<':
		if l.Peek() == '=' {
			l.Next()
			return LESS_EQUAL
		}
		return LESS
	case next == '>':
		if l.Peek() == '=' {
			l.Next()
			return GREATER_EQUAL
		}
		return GREATER
	case next == '{':
		return BEGIN_BLOCK
	case next == '}':
		return END_BLOCK
	case next == '+':
		return PLUS
	case next == '/':
		return DIVIDE
	case next == '*':
		return MULTI
	case next == '(':
		return int(next)
	case next == ')':
		return int(next)
	case next == '\n':
		l.scanWhitespace()
		if l.Peek() == ';' {
			l.Skip()
			return COLON
		}
		return COLON
	default:
		if unicode.IsSpace(next) {
			l.scanWhitespace()
			l.buffer.Reset()
			goto retry
		} else if next == eof {
			return eof
		} else if isDigit(next) {
			l.scanDigit(next)
			return NUMBER
		}
		l.scanIdentifier()
		text := l.TokenText()
		if text == "false" {
			return FALSE
		} else if text == "true" {
			return TRUE
		} else if text == "let" {
			return LET
		} else if text == "if" {
			return IF
		} else if text == "then" {
			return THEN
		} else if text == "else" {
			return ELSE
		}
		return ID
	}
}

func (l *Lexer) Next() rune {
	r, w, err := l.input.ReadRune()
	if err == io.EOF {
		return eof
	}
	if r == '\n' {
		l.position = &Position{line: l.position.line + 1}
	}
	l.position.column += w
	l.offset += w
	l.buffer.WriteRune(r)
	return r
}

func (l *Lexer) Skip() rune {
	r, w, err := l.input.ReadRune()
	if err == io.EOF {
		return eof
	}
	if r == '\n' {
		l.position = &Position{line: l.position.line + 1}
	}
	l.position.column += w
	l.offset += w
	return r
}

func (l *Lexer) Peek() rune {
	lead, err := l.input.Peek(1)
	if err == io.EOF {
		return eof
	} else if err != nil {
		l.Error(err.Error())
		return 0
	}

	p, err := l.input.Peek(runeLen(lead[0]))

	if err == io.EOF {
		return eof
	} else if err != nil {
		l.Error("unexpected input error")
		return 0
	}

	ruNe, _ := utf8.DecodeRune(p)
	return ruNe
}

func runeLen(lead byte) int {
	if lead < 0xC0 {
		return 1
	} else if lead < 0xE0 {
		return 2
	} else if lead < 0xF0 {
		return 3
	}
	return 4
}

// Lex Create Lexer
func (l *Lexer) Lex(lval *yySymType) int {
	if l.error != nil {
		return -1
	}
	typ := l.Scan()
	text := l.TokenText()
	lval.token = Token{typ: typ, literal: text}
	l.buffer.Reset()
	return typ
}
