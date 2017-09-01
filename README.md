# unparam

[![Build Status](https://travis-ci.org/mvdan/unparam.svg?branch=master)](https://travis-ci.org/mvdan/unparam)

	go get -u mvdan.cc/unparam

Reports unused function parameters and results in your code.

It also reports parameters that always receive the same values, results
that always return the same values and results that are never used.

To minimise false positives, it ignores:

* Unnamed and underscore parameters
* Funcs that may satisfy an interface
* Funcs that may satisfy a function signature
* Funcs that are stubs (empty, only error, only return constants, etc)
* Funcs that have multiple implementations via build tags

False positives can still occur by design. The aim of the tool is to be
as precise as possible - if you find any mistakes, file a bug.
