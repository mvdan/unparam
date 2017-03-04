package foo

import "net/http"

func AllUsed(a, b int) int {
	return a + b
}

func OneUnused(a, b int) int {
	return a
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi"))
}

type FooIface interface {
	foo(w http.ResponseWriter, code int) error
}

func FooImpl(w http.ResponseWriter, code int) error {
	w.Write([]byte("hi"))
	return nil
}

type FooType int

func (f FooType) AllUsed(a, b int) int {
	return a + b
}
