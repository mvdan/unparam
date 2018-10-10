package foo

func ClosureUse() {
	var enclosed FooType
	setValue := func(v *FooType) {
		enclosed = 2
	}
	var newValue FooType = 4
	println(enclosed)
	setValue(&newValue)
	println(enclosed)
}
