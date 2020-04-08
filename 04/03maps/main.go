package main

import "fmt"

func main() {
	idioms := make(map[string]string)
	idioms["en"] = "English"
	idioms["es"] = "Spanish"
	idioms["fr"] = "Franch"

	fmt.Println(idioms) // map[en:English es:Spanish fr:Franch]

	cities := map[string]string{"css": "Caracas", "ny": "New York", "bg": "Bogotá"}

	fmt.Println(cities)        // map[bg:Bogotá css:Caracas ny:New York]
	fmt.Println(cities["css"]) // Caracas

	delete(cities, "ny")

	fmt.Println(cities) // map[bg:Bogotá css:Caracas]

	if city, ok := cities["notexist"]; ok {
		fmt.Println("The value does exist and is", city)
	} else {
		fmt.Println("The value doesn't exist") // The value doesn't exist
	}

	numbers := map[int]string{
		1:    "hello",
		2016: "world",
		-999: "nice",
	}

	fmt.Println(numbers) // map[-999:nice 1:hello 2016:world]
}
