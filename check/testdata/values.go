package foo

import "math/rand"

func receivesSame(r rune) {
	doWork()
	if r == '0' {
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
	receivesSame('a')
	receivesSame('a')
	receivesDifferent('a')
	receivesDifferent('b')
	ReceivesSameExported('b')
	ReceivesSameExported('b')
	receivesCallExpr(randRune())
	receivesCallExpr(randRune())
}
