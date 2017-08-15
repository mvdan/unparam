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
