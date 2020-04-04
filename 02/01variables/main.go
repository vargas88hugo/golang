package main

import "fmt"

func main() {
	/*
		An example of regular declaration
		var firstName, lastName string
		firstName = "Hugo"
		lastName = "Vargas"
	*/

	// dynamic way to declare a variable and assign its
	// data type instantly
	firstName, lastName := "Hugo", "Vargas"

	firstName = "Mairelys"

	fmt.Println(firstName, lastName)

	// Regular way to declare a variable, first assigning
	// type and then value
	var a int
	var b int8

	a = 121212
	b = 5

	c := a + int(b)

	fmt.Printf("firstName %d lastName\n", c)
	fmt.Printf("b is the type %T\n", b)

	var name string
	var number int
	var boolean bool

	fmt.Println(name)
	fmt.Println(number)
	fmt.Println(boolean)
}

/*
Mairelys Vargas
firstName 121217 lastName
b is the type int8

0
false
*/
