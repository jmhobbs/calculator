%{
package calculator

func setResult(l yyLexer, r Result) {
  l.(*lexer).result = r
}
%}

%union{
  integer int
}

%token <integer> Integer

%token LexError

%token Addition
%token Subtraction
%token Multiplication
%token Divison

%start result

%%

result:
{
  setResult(yylex, Result{})
}
