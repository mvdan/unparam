package foo

func UsedInFuncLit(s string) func() {
	return func() {
		println(s)
	}
}
