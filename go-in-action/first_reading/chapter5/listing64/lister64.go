package main

import (
	"chapter5/counters"
	"fmt"
)

func main() {
	//counter := counters.alertCounter(10)
	counter := counters.New(10)
	fmt.Printf("카운터: %d\n", counter)
}
