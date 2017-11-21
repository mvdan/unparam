package foo

func allUsed(a, b FooType) FooType { return a + b }

func oneUnused(a, b FooType) FooType {
	FooSink = FooHolder{a}
	return 123
}

func structUnused(f FooStruct) {
	doWork()
}
