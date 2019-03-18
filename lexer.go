package calculator

import (
	"bytes"
	"errors"
	"strconv"
	"unicode"
)

//go:generate goyacc -c -cr -o parser.go parser.y

type Result struct {
	Operations []Operation
}

type Operation struct {
	Operand  int
	Operator byte
}

type lexer struct {
	result Result
	err    error
	input  []byte
	p      int
}

// Satisfies yyLexer interface
func (l *lexer) Error(s string) {
	l.err = errors.New(s)
}

// Satisfies yyLexer interface
func (l *lexer) Lex(lval *yySymType) int {
	for b := l.next(); b != 0; b = l.next() {
		switch {
		case unicode.IsSpace(rune(b)):
			continue
		case unicode.IsDigit(rune(b)):
			l.backup()
			return l.scanInteger(lval)
		case b == '+':
			return ADDITION
		case b == '-':
			return SUBTRACTION
		case b == '*':
			return MULTIPLICATION
		case b == '/':
			return DIVISION
		default:
			return int(b)
		}
	}
	return 0
}

func (l *lexer) scanInteger(lval *yySymType) int {
	buf := bytes.NewBuffer(nil)
	for b := l.next(); ; b = l.next() {
		switch {
		case unicode.IsDigit(rune(b)):
			buf.WriteByte(b)
		default:
			if b != 0 {
				l.backup()
			}
			i, err := strconv.Atoi(buf.String())
			if err != nil {
				return LexError
			}
			lval.int = i
			return Integer
		}
	}
}

// Gets the next byte in the input.
func (l *lexer) next() byte {
	if l.p >= len(l.input) {
		return 0
	}
	b := l.input[l.p]
	l.p++
	return b
}

// Seek input backwards one byte.
func (l *lexer) backup() {
	l.p--
	if l.p < 0 {
		l.p = 0
	}
}
