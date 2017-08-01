package foo

import "net/http"

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi"))
}

type FooIface interface {
	foo(w http.ResponseWriter, code FooType) error
}

type FooMayImpl struct{}

func (f FooMayImpl) foo(w http.ResponseWriter, code FooType) error {
	w.Write([]byte("hi"))
	return nil
}

func FooMayUse(f FooIface) {
	f.foo(nil, 0)
}
