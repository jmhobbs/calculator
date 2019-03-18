package calculator

//go:generate goyacc -c -cr -o parser.go parser.y

type Result struct{}

type lexer struct {
	result Result
}
