package foo

func (f FooType) AllUsed(a FooType) FooType { return f + a }

func (f FooType) OneUnused(a FooType) FooType { return 2 * f }
