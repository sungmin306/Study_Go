package main

import "fmt"

var array3 [5]string

func main() {
	array := [5]int{10, 20, 30, 40, 50}
	fmt.Print(array)
	array[2] = 35
	fmt.Print(array)

	array2 := [5]*int{0: new(int), 1: new(int)}
	fmt.Print(array2)
	*array2[0] = 10
	*array2[1] = 20
	//*array2[2] = 30 -> 초기화 안돼서 에러
	fmt.Print(array2)

	array4 := [5]string{"Blue", "Red", "Green", "Yellow", "Pink"}
	array3 = array4
	fmt.Println(array3)
	fmt.Println(array4)
}
