package calculator

//go:generate golex -o lexer.go lexer.l

//go:generate goyacc -c -cr -o parser.go parser.y

func Parse(input string) ([]Operation, error) {
	l := newLexer([]byte(input))
	_ = yyParse(l)
	if l.err != nil {
		return []Operation{}, l.err
	}
	return l.result.Operations, nil
}

type Result struct {
	Operations []Operation
}

type Operation struct {
	Operand  int
	Operator byte
}
