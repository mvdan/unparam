package foo

func singleReturn() int {
	doWork()
	return 3
}

func manyReturns() int {
	if cond {
		doWork()
		return 3
	}
	return 3
}

func manyReturnsDifferent(b bool) int {
	for cond {
		doWork()
		if b {
			return 4
		}
	}
	return 3
}

func manyReturnsMultiple() (b bool, s string) {
	if cond {
		doWork()
		return true, "foo"
	}
	return true, "foo"
}

func singleNilError() (bool, error) {
	doWork()
	return true, nil
}

func manyNilError() (bool, error) {
	if cond {
		doWork()
		return false, nil
	}
	return true, nil
}
