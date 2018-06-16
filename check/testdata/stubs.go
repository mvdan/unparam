package foo

import (
	"errors"
	"log"
	"net/http"
)

func dummyImpl(f FooType) {}

func panicImpl(f FooType) { panic("dummy") }

func nonPanicImpl(w http.ResponseWriter, f FooType) {
	for i := 0; i < 10; i++ {
		w.Write([]byte("foo"))
	}
	panic("default")
}

func throw(v ...interface{}) {}

func throwImpl(f FooType) { throw("dummy") }

func endlessLoop(w http.ResponseWriter) {
	for {
		w.Write([]byte("foo"))
	}
}

func nonPanicImpl2(w http.ResponseWriter, f FooType) {
	endlessLoop(w)
	panic("unreachable")
}

func zeroImpl(f FooType) (int, string, []byte) { return 0, "", nil }

func errorsImpl(f FooType) error { return errors.New("unimpl") }

const constFoo = FooType(123)

func (f FooType) Error() string { return "foo" }

func customErrImpl(f FooType) error { return constFoo }

func nonConstImpl(f FooType, s string) error { return Sink.(error) }

func logImpl(f FooType) { log.Print("not implemented") }

func oneOverwritten(a FooType, i uint8) (FooType, uint8) {
	i = 3
	a += 1
	return a, i
}

func zeroStructImpl(f FooStruct) FooStruct {
	return FooStruct{}
}

func zeroMapImpl(f FooStruct) map[int]int {
	return map[int]int{}
}

func zeroMapMakeImpl(f FooStruct) map[int]int {
	return make(map[int]int, 0)
}

func customError(msg string) error { return errors.New("custom: " + msg) }
func customErrorImpl(f FooType) (int8, error) {
	return -1, customError("x: bar")
}

func customErrorf(msg string) error { return errors.New("custom: " + msg) }
func customErrorfImpl(f FooType) (int8, error) {
	return -1, customErrorf("x: bar")
}
