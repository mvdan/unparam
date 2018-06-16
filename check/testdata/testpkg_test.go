package foo_test

type FooTypeTest uint

func oneUnused3(a, b FooTypeTest) FooTypeTest {
	a += 2
	return a
}
