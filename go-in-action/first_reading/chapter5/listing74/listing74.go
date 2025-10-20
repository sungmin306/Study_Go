package main

import (
	"chatper5/entities"
	"fmt"
)

func main() {
	a := entities.Admin{
		Rights: 10,
	}

	a.Name = "Bill"
	a.Email = "bill@email.com"

	fmt.Printf("사용자: %v\n", a)
}
