//+build !other

package foo

func OneUnusedMain(a, b FooType) FooType {
	return a + 2
}

func MultipleImpls(f FooType) int32 {
	doWork()
	return 3
}

func (f FooType) MultImplsMethod(f2 FooType) uint32 {
	doWork()
	return 3
}

func (f *FooType) MultImplsMethod2(f3 FooType) int64 {
	doWork()
	return 3
}

func MultImplsMethod(f FooType) uint64 {
	doWork()
	return 3
}
