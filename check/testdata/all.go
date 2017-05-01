package foo

import (
	"errors"
	"log"
	"net/http"
)

type FooType int

func AllUsed(a, b FooType) FooType { return a + b }

func OneUnused(a, b FooType) FooType { return a }

func doWork() {}

var Sink interface{}

func Parent() {
	oneUnused := func(f FooType) {
		doWork()
	}
	Sink = oneUnused
}

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

func (f FooType) AllUsed(a FooType) FooType { return f + a }

func (f FooType) OneUnused(a FooType) FooType { return f }

func DummyImpl(f FooType) {}

func PanicImpl(f FooType) { panic("dummy") }

func NonPanicImpl(w http.ResponseWriter, f FooType) {
	for i := 0; i < 10; i++ {
		w.Write([]byte("foo"))
	}
	panic("default")
}

func endlessLoop(w http.ResponseWriter) {
	for {
		w.Write([]byte("foo"))
	}
}

func NonPanicImpl2(w http.ResponseWriter, f FooType) {
	endlessLoop(w)
	panic("unreachable")
}

func throw(v ...interface{}) {}

func ThrowImpl(f FooType) { throw("dummy") }

func ZeroImpl(f FooType) (int, string, []byte) { return 0, "", nil }

func ErrorsImpl(f FooType) error { return errors.New("unimpl") }

const ConstFoo = FooType(123)

func (f FooType) Error() string { return "foo" }

func CustomErrImpl(f FooType) error { return ConstFoo }

func NonConstImpl(f FooType, s string) error { return f }

func LogImpl(f FooType) { log.Print("not implemented") }

type Foo2Func func(a FooType, s string) int

func Foo2Impl(a FooType, s string) int { return int(a) }

func NoName(FooType) { doWork() }

func UnderscoreName(_ FooType) { doWork() }

type BarStruct struct {
	Fn func(a FooType, b byte)
}

func BarField(a FooType, b byte) { doWork() }

type Bar2Struct struct {
	St struct {
		fn func(a FooType, r rune)
	}
}

func Bar2Field(a FooType, r rune) { doWork() }

func FuncAsParam(fn func(FooType) string) { fn(0) }

func PassedAsParam(f FooType) string {
	doWork()
	return "foo"
}

func (f FooType) FuncAsParam2(fn func(FooType) []byte) { fn(0) }

func PassedAsParam2(f FooType) []byte {
	doWork()
	return nil
}

type RecursiveIface interface {
	Foo(RecursiveIface)
}

func AsSliceElem(f FooType) []int {
	doWork()
	return nil
}

var SliceElems = []func(FooType) []int{AsSliceElem}

func AnonType() {
	for _, f := range []func(FooType, int32){
		func(f FooType, i int32) {
			doWork()
			println(i)
		},
	} {
		f(1, 2)
	}
	for _, f := range []struct {
		f2 func(f FooType, i int64)
	}{
		{f2: func(f FooType, i int64) {
			doWork()
			println(i)
		}},
	} {
		f.f2(3, 4)
	}
}

func UsedAsArg() {
	foo := func(f func(f FooType, u uint32)) {
		f(5, 6)
	}
	bar := func(v interface{}) {
		doWork()
		println(v)
	}
	foo(func(f FooType, u uint32) {
		println(f)
	})
	bar(func(f FooType, u uint64) {
		println(f)
	})
}

func globalParam(f func(f FooType, i int8)) {
	f(7, 8)
}

func UsedAsGlobalArg(f FooType, i int8) {
	doWork()
	println(f)
}

func globalParamIface(v interface{}) {
	println(v)
}

func UsedAsGlobalArgIface(f FooType, i int16) {
	doWork()
	println(f)
}

func GlobArgUse() {
	globalParam(UsedAsGlobalArg)
	globalParamIface(UsedAsGlobalArgIface)
}

func OneOverwritten(a FooType, i uint8) (FooType, uint8) {
	i = 3
	return a, i
}

type barIface interface {
	bar(FooType, uint16)
}

type barType struct{}

func (b *barType) bar(f FooType, u uint16) {
	doWork()
	println(f)
}

func barImpl() barIface { return &barType{} }

func BarIfaceUse() {
	b := barImpl()
	b.bar(0, 1)
}
