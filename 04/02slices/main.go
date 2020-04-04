package main

import "fmt"

/**
 * Slices
 * 1. it has a pointer
 * 2. it doesn't have a fixed size
 */
func main() {
	var names []string

	names = append(names, "Hugo")
	fmt.Printf("Size array is %d and capacity array is %d\n", len(names), cap(names))
	// Size array is 1 and capacity array is 1
	names = append(names, "Mairelys")
	fmt.Printf("Size array is %d and capacity array is %d\n", len(names), cap(names))
	// Size array is 2 and capacity array is 2
	names = append(names, "Elsa")
	fmt.Printf("Size array is %d and capacity array is %d\n", len(names), cap(names))
	// Size array is 3 and capacity array is 4
	names = append(names, "Daniel")
	fmt.Printf("Size array is %d and capacity array is %d\n", len(names), cap(names))
	// Size array is 4 and capacity array is 4
	names = append(names, "Nohelia")
	fmt.Printf("Size array is %d and capacity array is %d\n", len(names), cap(names))
	// Size array is 5 and capacity array is 8

	names[3] = "Victor"

	fmt.Println(names) // [Hugo Mairelys Elsa Victor Nohelia]

	// make(data type, initial size, initial capacity (optional))
	cities := make([]string, 0)

	cities = append(cities, "Bogota")
	cities = append(cities, "Caracas")
	cities = append(cities, "Miami")

	fmt.Println(cities) // [Bogota Caracas Miami]

	slice := []string{
		"one",
		"two",
		"three",
	}

	fmt.Println(slice) // [one two three]
}
