package foo

import "errors"

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

func singleIgnoredUnderscore() (_ rune) {
	doWork()
	return '0'
}

func SingleIgnoredUnderscoreUse() {
	singleIgnoredUnderscore()
	_ = singleIgnoredUnderscore()
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

func errorIgnored() (int, error) {
	doWork()
	if cond {
		return 3, errors.New("foo")
	}
	return 2, nil
}

func ErrorIgnoredUse() {
	errorIgnored()
	i, _ := errorIgnored()
	println(i)
}

func ignoredGoDefer() (int, string) {
	doWork()
	return 2, "bar"
}

func IgnoredGoDeferUse() {
	go ignoredGoDefer()
	defer ignoredGoDefer()
	i, _ := ignoredGoDefer()
	println(i)
}
