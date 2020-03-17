package main

import "fmt"

func main() {
	if 5 < 10 || 1 > 7 {
		fmt.Println("If block 01!")
	} else {
		fmt.Println("Else block 01")
	}

	if !(5 < 10 || 1 > 7) {
		fmt.Println("In the block 01!")
	} else if true {
		fmt.Println("Else If block 02")
	} else {
		fmt.Println("Else block 02")
	}

	fmt.Println("End of the program")
}
