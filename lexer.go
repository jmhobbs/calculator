package calculator

//go:generate goyacc -c -cr -o parser.go parser.y

type Result struct {
	Operations []Operation
}

type Operation struct {
	Operator byte
	Operand  int
}

type lexer struct {
	result Result
}
