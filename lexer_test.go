package calculator

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestLexer(t *testing.T) {
	cases := []struct {
		input  string
		tokens []int
	}{
		{
			"8 600",
			[]int{Integer, Integer},
		},
		{
			"5+6",
			[]int{Integer, ADDITION, Integer},
		},
		{
			"500-60",
			[]int{Integer, SUBTRACTION, Integer},
		},
		{
			"80*5",
			[]int{Integer, MULTIPLICATION, Integer},
		},
		{
			"90/10",
			[]int{Integer, DIVISION, Integer},
		},
	}

	for _, tc := range cases {
		var (
			lval   yySymType
			token  int
			tokens []int
		)

		l := newLexer([]byte(tc.input))
		for {
			token = l.Lex(&lval)
			if token == 0 {
				break
			}
			tokens = append(tokens, token)
			if token == LexError {
				break
			}
		}

		if !reflect.DeepEqual(tokens, tc.tokens) {
			t.Errorf("Unexpected tokens for %q\n  Expect: %v\n     Got: %v", tc.input, tokensToString(tc.tokens), tokensToString(tokens))
			return
		}
	}
}

func tokensToString(tkns []int) string {
	strs := []string{}
	for _, token := range tkns {
		switch token {
		case LexError:
			strs = append(strs, "LexError")
		case Integer:
			strs = append(strs, "Integer")
		case ADDITION:
			strs = append(strs, "ADDITION")
		case SUBTRACTION:
			strs = append(strs, "SUBTRACTION")
		case MULTIPLICATION:
			strs = append(strs, "MULTIPLICATION")
		case DIVISION:
			strs = append(strs, "DIVISION")
		default:
			if token < 128 {
				strs = append(strs, string(token))
			} else {
				strs = append(strs, fmt.Sprintf("UNKNOWN (%d)", token))
			}
		}
	}
	return strings.Join(strs, ", ")
}
