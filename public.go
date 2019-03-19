package calculator

func Parse(input string) ([]Operation, error) {
	l := &lexer{input: []byte(input)}
	_ = yyParse(l)
	if l.err != nil {
		return []Operation{}, l.err
	}
	return l.result.Operations, nil
}
