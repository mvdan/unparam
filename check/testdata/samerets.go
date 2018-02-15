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

func manyReturnsForwarded(r rune) int {
	if r == '3' {
		return 5
	}
	doWork()
	return 5
}

func forwarding(r rune) int {
	doWork()
	return manyReturnsForwarded(r)
}

func doubleReturnForwarded() (int, error) {
	doWork()
	return 5, nil
}

func forwardingDouble() (int, error) {
	doWork()
	return doubleReturnForwarded()
}

func doubleReturnNotForwarded() (int, error) {
	doWork()
	return 5, nil
}

func falseForwardinggDouble() (int, error) {
	doWork()
	n, err := doubleReturnNotForwarded()
	println(n)
	return 2, err
}

func neverForwarded() (int, error) {
	doWork()
	return 5, nil
}

func (f FooType) neverForwardedPtrMethod() (int, error) {
	doWork()
	return 5, nil
}

func neverForwarding() {
	doWork()
	neverForwarded()
	f := new(FooType)
	f.neverForwardedPtrMethod()
}
