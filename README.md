[![Build Status](https://travis-ci.org/jmhobbs/calculator.svg?branch=master)](https://travis-ci.org/jmhobbs/calculator) [![codecov](https://codecov.io/gh/jmhobbs/calculator/branch/master/graph/badge.svg)](https://codecov.io/gh/jmhobbs/calculator) [![GoDoc](https://godoc.org/github.com/jmhobbs/calculator?status.svg)](https://godoc.org/github.com/jmhobbs/calculator)

# Calculator

This is a demo of using goyacc and golex to build a simple parser which handles basic arithmetic.

It's a dumb little thing though, as it has no concept of the usual order of operations, it just goes left to right.

For example:

    jmhobbs@hobbs:cmd $ go run main.go 5+6*10
    5 + 6 * 10 = 110
    jmhobbs@hobbs:cmd $ go run main.go 10*5+6
    10 * 5 + 6 = 56
    jmhobbs@hobbs:cmd $ # Oh my sweet, simple calculator.

It was created as an example for a talk at the [Omaha Gophers](https://meetup.com/Omaha-Gophers) in March 2019.

You can view the (probably unhelpful) slides here: https://docs.google.com/presentation/d/1tRmOXhBPfYbEXXmD7qoRHYBNbjpQQSODPS5snRtm3Zo/edit

# Tags!

There are 5 tags on this repo which correspond to a state of completion.

  - [step-0](https://github.com/jmhobbs/calculator/tree/step-0) : Stub things out.
  - [step-1](https://github.com/jmhobbs/calculator/tree/step-1) : Define tokens.
  - [step-2](https://github.com/jmhobbs/calculator/tree/step-2) : Put together a parser.
  - [step-3](https://github.com/jmhobbs/calculator/tree/step-3) : Hand rolled lexer.
  - [step-4](https://github.com/jmhobbs/calculator/tree/step-4) : It works!
  - [step-5](https://github.com/jmhobbs/calculator/tree/step-5) : Build golex lexer.

# Links

  - [How to Write a Parser in Go](https://www.youtube.com/watch?v=NG0s3-s3whY)
  - [goyacc](https://godoc.org/modernc.org/goyacc)
  - [golex](https://godoc.org/modernc.org/golex)
