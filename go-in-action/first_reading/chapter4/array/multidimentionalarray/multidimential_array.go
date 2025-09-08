package main

import "fmt"

var array1 [4][2]int

func main() {
	array2 := [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	fmt.Println(array2)

	array3 := [4][2]int{1: {20, 21}, 3: {40, 41}}
	fmt.Println(array3)

	array4 := [4][2]int{1: {0: 20}, 3: {1: 4}}
	fmt.Println(array4)

	var array5 [2][2]int
	var array6 [2][2]int
	array5[0][0] = 10
	array5[0][1] = 20
	array5[1][0] = 30
	array5[1][1] = 40
	array6 = array5
	fmt.Println("array5 = ", array5)
	fmt.Println("array6 = ", array6)
	var array7 [2]int = array5[1]
	var value int = array5[1][0]
	fmt.Println(array7)
	fmt.Println(value)
}
