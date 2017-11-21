package foo

func allUsed(a, b FooType) FooType { return a + b }

func oneUnused(a, b FooType) FooType {
	a += 1
	return a
}

func structUnused(f FooStruct) {
	doWork()
}
