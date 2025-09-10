package main

import "fmt"

func main() {
	slice := make([]string, 5) // 길이만 설정하면 최대용량도 같은 값으로 설정됨
	fmt.Println(slice)

	slice2 := make([]int, 3, 5) // 길이 3, 최대 용량 5
	fmt.Println(slice2)

	slice3 := []string{"Red", "Blue", "Green", "Yellow", "Pink"}
	fmt.Println(slice3)

	slice4 := []int{10, 20, 30}
	fmt.Println(slice4)

	slice5 := []string{99: ""}
	fmt.Println(slice5)

	array := [3]int{10, 20, 30}
	fmt.Println(array)
	slice6 := []int{10, 20, 30}
	fmt.Println(slice6)

	//var slice7 []int // nil 슬라이스 생성
	slice8 := make([]int, 0)
	fmt.Println(slice8)

	slice9 := []int{}
	fmt.Println(slice9)
}
