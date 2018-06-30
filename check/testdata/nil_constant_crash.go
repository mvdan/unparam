package foo

import "log"

func f(p *int) {
	if p != nil {
		log.Printf("p is %p", p)
	}
}

func a() {
	f(nil)
}

func b() {
	f(nil)
}

func c() {
	f(nil)
}

func d() {
	f(nil)
}
