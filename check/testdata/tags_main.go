//+build !other

package foo

func oneUnusedMain(a, b FooType) FooType {
	return a + 2
}

func multipleImpls(f FooType) int32 {
	doWork()
	return 3
}

func (f FooType) multImplsMethod(f2 FooType) uint32 {
	doWork()
	return 3
}

func (f *FooType) multImplsMethod2(f3 FooType) int64 {
	doWork()
	return 3
}

func multImplsMethod(f FooType) uint64 {
	doWork()
	return 3
}
