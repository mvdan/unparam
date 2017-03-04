package foo

import "net/http"

func allUsed(a, b int) int {
	return a + b
}

func oneUnused(a, b int) int {
	return a
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi"))
}

type fooIface interface {
	foo(w http.ResponseWriter, code int) error
}

func fooImpl(w http.ResponseWriter, code int) error {
	w.Write([]byte("hi"))
	return nil
}
