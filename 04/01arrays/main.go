package main

import "fmt"

/**
 * All the values in an array must be the same data type
 * The length of an array is fixed
 */
func main() {
	var names [3]string
	names[0] = "Hugo"
	names[1] = "Mairelys"
	names[2] = "Elsa"

	fmt.Println(names) // [Hugo Mairelys Elsa]

	cities := [3]string{"Bogota", "New York", "Miami"}
	city := cities[1]

	fmt.Println(cities) // [Bogota New York Miami]
	fmt.Println(city)   // New York

	size := len(cities)

	fmt.Println("The size of array is:", size) // The size of array is 3

	fmt.Println(cities[0:2]) // [Bogota New York]

	array := [3]string{
		"one",
		"two",
		"three", // the last element must have a comma if the array have new lines
	}

	fmt.Println(array) // [one two three]
}
