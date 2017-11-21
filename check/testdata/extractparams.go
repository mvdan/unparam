package foo

func funcResultAsParams(a, b FooType) FooType {
	doWork()
	return a + 1
}

func (f FooType) methodResultAsParams(a, b FooType) FooType {
	doWork()
	return a + 1
}

func generateResults() (a, b FooType) { return }

func funcResultAsParam(a FooType) FooType {
	doWork()
	return 3
}

func (f FooType) methodResultAsParam(a FooType) FooType {
	doWork()
	return 3
}

func generateResult() (a FooType) { return }

func funcResultsAsParams(a, b FooType) FooType {
	doWork()
	return a + 1
}

func ResultCalls() {
	var f FooType
	Sink = funcResultAsParams(1, 2)
	Sink = funcResultAsParams(generateResults())
	Sink = f.methodResultAsParams(1, 2)
	Sink = f.methodResultAsParams(generateResults())
	Sink = funcResultAsParam(4)
	Sink = funcResultAsParam(generateResult())
	Sink = f.methodResultAsParam(4)
	Sink = f.methodResultAsParam(generateResult())
	Sink = funcResultsAsParams(generateResult(), generateResult())
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
	_, Sink = returnResultsOwn()
	_, Sink = returnResultsDirectly()
	_, Sink = returnResultsDirectly()
}
