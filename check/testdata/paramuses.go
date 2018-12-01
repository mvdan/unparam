package foo

import (
	"net/http"
	"time"
)

func UsedInFuncLit(s string) func() {
	return func() {
		println(s)
	}
}

func StructUsedInField(path string, expiry time.Time) {
	Sink = http.Cookie{Path: path, Expires: expiry}
}

func usedInUnderscoreAssignment(s string, d time.Duration) {
	_ = s
	doWork()
	Sink = d
}

func notUsedInUnderscoreAssignment(s string, d time.Duration) {
	if d > 0 {
		s := 123
		_ = s
	}
	doWork()
	Sink = d
}
