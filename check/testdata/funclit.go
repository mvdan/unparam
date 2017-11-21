package foo

func parent() {
	oneUnused := func(f FooType) {
		doWork()
	}
	Sink = oneUnused
}
