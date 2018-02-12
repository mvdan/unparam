package main

func OneUnused(a, b uint) uint {
	a += 1
	return a
}

type fnType func(a, b int) int

func mightImplement(a, b int) int {
	a *= 4
	return a
}

func main() {
	var f fnType
	f(2, 3)
}
