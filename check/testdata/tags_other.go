//+build other

package foo

func multipleImpls(f FooType) int32 {
	doWork()
	return int32(f)
}

func (f FooType) multImplsMethod(f2 FooType) uint32 {
	doWork()
	return uint32(f)
}

func (f *FooType) multImplsMethod2(f3 FooType) int64 {
	doWork()
	return int64(*f)
}
