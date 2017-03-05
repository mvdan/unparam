package foo

import "net/http"

type FooType int

func AllUsed(a, b FooType) FooType { return a + b }

func OneUnused(a, b FooType) FooType { return a }

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi"))
}

type FooIface interface {
	foo(w http.ResponseWriter, code FooType) error
}

func FooImpl(w http.ResponseWriter, code FooType) error {
	w.Write([]byte("hi"))
	return nil
}

func (f FooType) AllUsed(a, b FooType) FooType { return a + b }

func DummyImpl(f FooType) {}

func PanicImpl(f FooType) { panic("dummy") }

func NonPanicImpl(w http.ResponseWriter, f FooType) {
	for i := 0; i < 10; i++ {
		w.Write([]byte("foo"))
	}
	panic("dummy")
}

func throw(v ...interface{}) {}

func ThrowImpl(f FooType) { throw("dummy") }

func ZeroImpl(f FooType) (int, string, []byte) { return 0, "", nil }

type BarFunc func(a FooType, s string) int

func BarImpl(a FooType, s string) int { return int(a) }

func NoName(FooType) { println("foo") }

func UnderscoreName(_ FooType) { println("foo") }

type BarStruct struct {
	fn func(a FooType, b byte)
}

func BarField(a FooType, b byte) { println(a) }

func FuncAsParameter(fn func(FooType) string) { fn(0) }

func PassedAsParam(f FooType) string { return "foo" }
