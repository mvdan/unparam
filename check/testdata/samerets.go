package foo

func SingleReturn() int {
	doWork()
	return 3
}

func ManyReturns() int {
	if cond {
		doWork()
		return 3
	}
	return 3
}

func ManyReturnsDifferent(b bool) int {
	for cond {
		doWork()
		if b {
			return 4
		}
	}
	return 3
}

func ManyReturnsMultiple() (b bool, s string) {
	if cond {
		doWork()
		return true, "foo"
	}
	return true, "foo"
}
