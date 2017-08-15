package foo

func singleIgnored() rune {
	doWork()
	return '0'
}

func SingleIgnoredUse() {
	singleIgnored()
	_ = singleIgnored()
}

func singleNotIgnored() rune {
	doWork()
	return '0'
}

func SingleNotIgnoredUse() {
	singleNotIgnored()
	r := singleNotIgnored()
	println(r)
}

func singleIgnoredName() (r rune) {
	doWork()
	return '0'
}

func SingleIgnoredNameUse() {
	singleIgnoredName()
	_ = singleIgnoredName()
}

func allIgnored() (int, string) {
	doWork()
	return 2, "foo"
}

func AllIgnoredUse() {
	allIgnored()
	_, _ = allIgnored()
}

func someIgnored() (int, string) {
	doWork()
	return 2, "foo"
}

func SomeIgnoredUse() {
	someIgnored()
	i, _ := someIgnored()
	println(i)
}
