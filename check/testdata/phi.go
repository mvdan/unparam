package foo

func allUsedPhi(a, b FooType) *FooType {
	doWork()
	c := a + b
	return &c
}

func oneUnusedPhi(a, b FooType) *FooType {
	doWork()
	c := a + 4
	return &c
}

func PhiUse(fn2 bool) {
	var a, b FooType = 2, 3
	fn := allUsedPhi
	if fn2 {
		fn = oneUnusedPhi
	}
	println(fn(a, b))
}
