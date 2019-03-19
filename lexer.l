%{
package calculator

import (
	"errors"
	"log"
	"strconv"
)

type lexer struct {
	result   Result
	err      error
	input    []byte
	p        int
	buf      []byte
	current  byte
	previous byte
}

func newLexer(input []byte) *lexer {
	l := &lexer{input: input}
	l.next() // Prime the pump
	return l
}

// Gets the next byte in the input.
func (l *lexer) next() byte {
	// Add the current byte to the buffer if set.
	if l.current != 0 {
		l.buf = append(l.buf, l.current)
		l.previous = l.current
	}
	// Reset current
	l.current = 0
	// If we're out of bytes, set to 0.
	if l.p >= len(l.input) {
		l.current = 0
	// Otherwise, get the next byte and increment input index
	} else {
		l.current = l.input[l.p]
		l.p++
	}
	// Here's the byte you ordered!
	return l.current
}

// Error satisfies yyLexer.
func (l *lexer) Error(s string) {
	l.err = errors.New(s)
}

//    %yyc is a "macro" to access the "current" character.
//
//    %yyn is a "macro" to move to the "next" character.
//
//    %yyb is a "macro" to return the beginning-of-line status (a bool typed value).
//        It is used for patterns like `^re`.
//        Example: %yyb prev == 0 || prev == '\n'
//
//    %yyt is a "macro" to return the top/current start condition (an int typed value).
//        It is used when there are patterns with conditions like `<cond>re`.
//        Example: %yyt startCond

// Lex satisfies yyLexer.
func (l *lexer) Lex(lval *yySymType) int {
%}

%yyc l.current
%yyn l.next()

D   [0-9]+

%%
		// Code before the first rule is executed before every scan cycle (state 0 action)
		l.buf = l.buf[:0] // Truncate the input buffer.

[ \t\n\r]+          // Ignore whitespace

"+"
		return ADDITION

"-"
		return SUBTRACTION

"*"
		return MULTIPLICATION

"/"
		return DIVISION

{D}
		i, err := strconv.Atoi(string(l.buf))
		if err != nil {
			return LexError
		}
		lval.int = i
		return Integer

\0
		return 0

.
		return int(l.current)

%%
		// The rendered scanner enters top of the user code section when
		// lexem recognition fails. In this example it should never happen.
		log.Println("scanner internal error")
		return LexError

} // this closes lexer.Lex(*yySymType) int
