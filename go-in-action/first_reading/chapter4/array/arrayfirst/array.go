package main

import "fmt"

func main() {
	var array [5]int                       // 배열 선언
	array2 := [5]int{10, 20, 30, 40, 50}   // 배열 리터럴로 선언
	array3 := [...]int{10, 20, 30, 40, 50} // 배열 길이 ... 지정
	array4 := [5]int{1: 10, 2: 20}         // 특정 index 숫자 넣기]

	fmt.Println(array)
	fmt.Println(array2)
	fmt.Println(array3)
	fmt.Println(array4)
	fmt.Println(1)
}
