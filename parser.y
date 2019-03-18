%{
package calculator

func setResult(l yyLexer, r Result) {
  l.(*lexer).result = r
}
%}

%union{
  int int
  byte byte
  operation Operation
  operations []Operation
}

%token <int> Integer

%token LexError

%token ADDITION
%token SUBTRACTION
%token MULTIPLICATION
%token DIVISION

%type <operation> operation
%type <operations> operations
%type <byte> operator ADDITION SUBTRACTION MULTIPLICATION DIVISION

%start result

%%

result: operations
{
  setResult(yylex, Result{$1})
}

operations: operation
{
  $$ = []Operation{$1}
}
          | operations operation
{
  $1 = append($1, $2)
  $$ = $1
}

operation: operator Integer
{
  $$ = Operation{$2, $1}
}
         | Integer
{
  $$ = Operation{$1, 0}
}

operator: ADDITION
{
  $$ = '+'
}
        | SUBTRACTION
{
  $$ = '-'
}
        | MULTIPLICATION
{
  $$ = '*'
}
        | DIVISION
{
  $$ = '/'
}
