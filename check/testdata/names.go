package foo

func regularName(f FooType) { doWork() }

func noName(FooType) { doWork() }

func underscoreName(_ FooType) { doWork() }

func zeroSizeStruct(f struct{}) { doWork() }

func zeroSizeArray(f [0]bool) { doWork() }
