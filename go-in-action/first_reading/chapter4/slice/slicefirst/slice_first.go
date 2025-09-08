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

}
