# unparam

	go get -u github.com/mvdan/unparam

Reports unused function parameters in your code.

To minimise false positives, it ignores:

* Unnamed and underscore parameters
* Funcs whose signature matches a top-level func type
* Funcs whose signature matches a top-level interface method
* Funcs that have empty bodies
* Funcs that will simply panic or return constants

Note that false positives can still occur by design. If you find any,
please file a bug. You can also use an underscore name for the parameter
to make the tool ignore it.
