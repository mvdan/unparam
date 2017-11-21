package foo

func reusedRecursively(f FooType, b bool) rune {
	if !b {
		return 0
	}
	return reusedRecursively(f, b)
}

func reusedRecursivelySwapped(f1, f2 FooType, b bool) rune {
	if !b {
		return 0
	}
	return reusedRecursivelySwapped(f2, f1, b)
}

func reusedRecursivelyModified(f FooType, b bool) rune {
	if !b {
		return 0
	}
	return reusedRecursivelyModified(f+FooType(1), b)
}

func unusedVariadic(a FooType, bs ...byte) {
	doWork()
	println(a)
}

func reusedRecursivelyVariadic(a FooType, bs ...byte) {
	if a == 0 {
		reusedRecursivelyVariadic(a, bs...)
	}
}
