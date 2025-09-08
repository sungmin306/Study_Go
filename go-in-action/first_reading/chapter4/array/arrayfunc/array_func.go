package main

import "fmt"

var array [1e6]int

func foo(array [1e6]int) {
	fmt.Println(array)
}

func foo2(array *[1e6]int) {
	fmt.Println(array)
}
func main() {
	//fmt.Println(array)
	foo(array)   // 함수에 100만크기의 배열이 복사!해서 넘어가지니 비효율적
	foo2(&array) // 주소값만 보내기에 효율적 but 성능 개선 여지 있음
}
