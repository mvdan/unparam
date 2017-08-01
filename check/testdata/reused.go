package foo

func ReusedRecursively(f FooType, b bool) rune {
	if !b {
		return 0
	}
	return ReusedRecursively(f, b)
}

func ReusedRecursivelySwapped(f1, f2 FooType, b bool) rune {
	if !b {
		return 0
	}
	return ReusedRecursivelySwapped(f2, f1, b)
}

func ReusedRecursivelyModified(f FooType, b bool) rune {
	if !b {
		return 0
	}
	return ReusedRecursivelyModified(f+FooType(1), b)
}

func UnusedVariadic(a FooType, bs ...byte) {
	doWork()
	println(a)
}

func ReusedRecursivelyVariadic(a FooType, bs ...byte) {
	if a == 0 {
		ReusedRecursivelyVariadic(a, bs...)
	}
}
