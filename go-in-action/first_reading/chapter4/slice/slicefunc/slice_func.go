package main

func foo(slice []int) []int {
	return slice
}
func main() {
	slice := make([]int, 1e6)
	slice = foo(slice)
}
