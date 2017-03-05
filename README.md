# unparam

	go get -u github.com/mvdan/unparam

Reports unused function parameters in your code.

To avoid false positives, it ignores:

* Unnamed and underscore parameters
* Funcs whose signature matches a top-level func type
* Funcs whose signature matches a top-level interface method
* Funcs that have empty bodies or simply panic
