package foo

type FooType int

type FooStruct struct {
	field int
}

func doWork() {}

var Sink interface{}

type FooHolder struct {
	f FooType
}

var FooSink FooHolder

var cond bool

const FooConst FooType = 123
