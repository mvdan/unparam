package foo

func RegularName(f FooType) { doWork() }

func NoName(FooType) { doWork() }

func UnderscoreName(_ FooType) { doWork() }

func ZeroSizeStruct(f struct{}) { doWork() }

func ZeroSizeArray(f [0]bool) { doWork() }
