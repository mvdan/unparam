package foo

import "math/rand"

func receivesSameMany(f FooType) {
	doWork()
	if f == 0 {
		println(f)
	}
}

func receivesSameManyLit(r rune) {
	doWork()
	if r == '0' {
		println(r)
	}
}

func receivesSameManyNamed(f FooType) {
	doWork()
	if f == 0 {
		println(f)
	}
}

func receivesSameManyMixed(f FooType) {
	doWork()
	if f == 1 {
		println(f)
	}
}

func receivesSameOnce(r rune) {
	doWork()
	if r == '1' {
		println(r)
	}
}

func receivesDifferent(r rune) {
	doWork()
	if r == '0' {
		println(r)
	}
}

func ReceivesSameExported(r rune) {
	doWork()
	if r == '0' {
		println(r)
	}
}

func receivesCallExpr(r rune) {
	doWork()
	if r == '0' {
		println(r)
	}
}

func randRune() rune { return rune(rand.Int31()) }

func withVariadic(s ...string) {
	doWork()
	println(len(s))
}

func receivesSameFromGenerated(f FooType) {
	doWork()
	if f == 4 {
		println(f)
	}
}

func CallReceivers() {
	receivesSameMany(3)
	receivesSameMany(3)
	receivesSameMany(3)
	receivesSameMany(3)
	receivesSameManyLit('a')
	receivesSameManyLit('a')
	receivesSameManyLit('a')
	receivesSameManyLit('a')
	receivesSameManyNamed(FooConst)
	receivesSameManyNamed(FooConst)
	receivesSameManyNamed(FooConst)
	receivesSameManyNamed(FooConst)
	receivesSameManyMixed(FooConst)
	receivesSameManyMixed(FooConst)
	receivesSameManyMixed(123)
	receivesSameManyMixed(FooType(123))
	receivesSameOnce('b')
	receivesDifferent('a')
	receivesDifferent('b')
	receivesDifferent('c')
	receivesDifferent('d')
	ReceivesSameExported('b')
	ReceivesSameExported('b')
	ReceivesSameExported('b')
	ReceivesSameExported('b')
	receivesCallExpr(randRune())
	receivesCallExpr(randRune())
	receivesCallExpr(randRune())
	receivesCallExpr(randRune())
	withVariadic()
}
