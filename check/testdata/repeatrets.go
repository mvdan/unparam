package foo

func ReturnParamUnchanged(n int) int {
	doWork()
	return n
}

func (f FooType) ReturnParamUnchanged(n int) int {
	doWork()
	return n
}

func ReturnParamChanged(n int) int {
	doWork()
	n += 2
	return n
}

func (f FooType) ReturnReceivingParamUnchanged() FooType {
	doWork()
	return f
}

func ReturnParamSliceUnchanged(ns []int) []int {
	doWork()
	return ns
}

func ReturnParamSliceChanged(ns []int) []int {
	doWork()
	ns = append(ns, 3)
	return ns
}

func ReturnParamPointerUnchanged(n *FooStruct) *FooStruct {
	doWork()
	return n
}

func ReturnParamPointerChanged(n *FooStruct) *FooStruct {
	doWork()
	println(n)
	n = nil
	return n
}

func ReturnParamPointerUnchanged2(n *FooStruct) *FooStruct {
	doWork()
	n.field = 3
	return n
}

func AppendImplStub(fs []FooStruct) []FooStruct { return fs }
