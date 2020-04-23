// Echo1 prints its command line arguments
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i] // Concatenates argument by argument
		sep = " "             // For separating each argument
	}

	fmt.Println(s)
}
