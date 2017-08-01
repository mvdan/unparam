package foo

func AllUsed(a, b FooType) FooType { return a + b }

func OneUnused(a, b FooType) FooType { return a }

func StructUnused(f FooStruct) {
	doWork()
}
