package main

import "fmt"

var colors2 map[string]string

func main() {
	dick := make(map[string]int)
	fmt.Println(dick)
	dick2 := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}
	fmt.Println(dick2)

	colors := map[string]string{}
	colors["Red"] = "#da1337"

	value, exists := colors["Blue"]
	if exists {
		fmt.Println(value)
	}

	value2 := colors["Blue"]
	if value2 != "" {
		fmt.Println(value2)
	}

	colors3 := map[string]string{
		"AliceBlue":   "#f0f8ff",
		"Coral":       "#ff7F50",
		"DarkGray":    "#a9a9a9",
		"ForestGreen": "#228b22",
	}

	for key, value := range colors3 {
		fmt.Printf("키 : %s 값: %s\n", key, value)
	}

	delete(colors3, "Coral")
	for key, value := range colors3 {
		fmt.Printf("키 : %s 값: %s\n", key, value)
	}
}
