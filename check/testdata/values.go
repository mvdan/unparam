package foo

import "math/rand"

func receivesSameMany(r rune) {
	doWork()
	if r == '0' {
		println(r)
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

func CallReceivers() {
	receivesSameMany('a')
	receivesSameMany('a')
	receivesSameMany('a')
	receivesSameMany('a')
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
}
