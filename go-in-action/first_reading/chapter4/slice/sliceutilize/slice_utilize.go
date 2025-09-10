package main

import "fmt"

func main() {
	// slice := []int{10, 20, 330, 40, 50}
	// newSlice := slice[1:3]
	// fmt.Println(newSlice)

	// slice := []int{10, 20, 30, 40, 50}
	// newSlice := slice[1:3]
	// newSlice = append(newSlice, 60)
	// fmt.Println(slice[3])

	// slice := []int{10, 20, 30, 40}
	// newSlice := append(slice, 50)
	// fmt.Println(newSlice)
	// source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	// slice := source[2:3:3]
	// slice = append(slice,"Kiwi")
	// s1 := []int{1, 2}
	// s2 := []int{3, 4}
	// fmt.Printf("%v\n", append(s1, s2...))
	// slice := []int{10, 20, 30, 40}

	// for index, value := range slice {
	// 	fmt.Printf("인덱스 %d 값 : %d\n", index, value)
	// }

	// slice := []int{10, 20, 30, 40}
	// for index, value := range slice {
	// 	fmt.Printf("값: %d, 값의 주소 : %X 원소의 주소 : %X\n", value, &value, &slice[index])
	// }

	// slice := []int{10, 20, 30, 40}
	// for _, value := range slice {
	// 	fmt.Printf("값 : %d\n", value)
	// }

	slice := []int{10, 20, 30, 40}

	for index := 2; index < len(slice); index++ {
		fmt.Printf("인덱스 : %d 값: %d\n", index, slice[index])
	}

}
