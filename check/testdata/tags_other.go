//+build other

package foo

func MultipleImpls(f FooType) int32 {
	doWork()
	return int32(f)
}

func (f FooType) MultImplsMethod(f2 FooType) uint32 {
	doWork()
	return uint32(f)
}

func (f *FooType) MultImplsMethod2(f3 FooType) int64 {
	doWork()
	return int64(*f)
}
