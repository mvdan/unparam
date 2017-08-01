package foo

import (
	"errors"
	"log"
	"net/http"
)

func DummyImpl(f FooType) {}

func PanicImpl(f FooType) { panic("dummy") }

func NonPanicImpl(w http.ResponseWriter, f FooType) {
	for i := 0; i < 10; i++ {
		w.Write([]byte("foo"))
	}
	panic("default")
}

func throw(v ...interface{}) {}

func ThrowImpl(f FooType) { throw("dummy") }

func endlessLoop(w http.ResponseWriter) {
	for {
		w.Write([]byte("foo"))
	}
}

func NonPanicImpl2(w http.ResponseWriter, f FooType) {
	endlessLoop(w)
	panic("unreachable")
}

func ZeroImpl(f FooType) (int, string, []byte) { return 0, "", nil }

func ErrorsImpl(f FooType) error { return errors.New("unimpl") }

const ConstFoo = FooType(123)

func (f FooType) Error() string { return "foo" }

func CustomErrImpl(f FooType) error { return ConstFoo }

func NonConstImpl(f FooType, s string) error { return f }

func LogImpl(f FooType) { log.Print("not implemented") }

func OneOverwritten(a FooType, i uint8) (FooType, uint8) {
	i = 3
	return a, i
}
