package foo

func FuncResultAsParams(a, b FooType) FooType {
	doWork()
	return a + 1
}

func (f FooType) MethodResultAsParams(a, b FooType) FooType {
	doWork()
	return a + 1
}

func generateResults() (a, b FooType) { return }

func FuncResultAsParam(a FooType) FooType {
	doWork()
	return 3
}

func (f FooType) MethodResultAsParam(a FooType) FooType {
	doWork()
	return 3
}

func generateResult() (a FooType) { return }

func FuncResultsAsParams(a, b FooType) FooType {
	doWork()
	return a + 1
}

func ResultCalls() {
	var f FooType
	Sink = FuncResultAsParams(1, 2)
	Sink = FuncResultAsParams(generateResults())
	Sink = f.MethodResultAsParams(1, 2)
	Sink = f.MethodResultAsParams(generateResults())
	Sink = FuncResultAsParam(4)
	Sink = FuncResultAsParam(generateResult())
	Sink = f.MethodResultAsParam(4)
	Sink = f.MethodResultAsParam(generateResult())
	Sink = FuncResultsAsParams(generateResult(), generateResult())
}

func returnResultsOwn() (FooType, FooType) {
	a := generateResult()
	a++
	return a, a * 2
}

func returnResultsDirectly() (FooType, FooType) {
	return generateResults()
}

func ReturnResultsCalls() {
	_, Sink = returnResultsOwn()
	_, Sink = returnResultsDirectly()
}
