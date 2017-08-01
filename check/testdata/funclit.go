package foo

func Parent() {
	oneUnused := func(f FooType) {
		doWork()
	}
	Sink = oneUnused
}
