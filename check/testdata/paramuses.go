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
