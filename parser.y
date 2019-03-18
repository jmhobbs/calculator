%{
package calculator

func setResult(l yyLexer, r Result) {
  l.(*lexer).result = r
}
%}

%union{
}

%start result

%%

result:
{
  setResult(yylex, Result{})
}
