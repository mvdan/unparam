package foo

func (f FooType) allUsed(a FooType) FooType { return f + a }

func (f FooType) oneUnused(a FooType) FooType { return 2 * f }
