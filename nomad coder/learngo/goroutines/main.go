package main

import (
	"fmt"
)

func main() {
	c := make(chan string)
	people := [5]string{"nico", "flynn", "dal", "japanguy", "larry"}
	for _, person := range people {
		go isSexy(person, c)
	}
	for i := 0; i < len(people); i++ {
		fmt.Println("waiting for", i)
		fmt.Println(<-c) // blocking operation
	}

}

func isSexy(person string, c chan string) {
	c <- person + " is sexy"
}
