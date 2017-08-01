package foo

type RecursiveIface interface {
	Foo(RecursiveIface)
}
