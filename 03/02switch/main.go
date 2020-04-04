package main

import "fmt"

func main() {
	var a int
	a = 2

	switch a {
	case 1:
		fmt.Println("The value is 1")
		break
	case 2:
		fmt.Println("The value is 2")
		break
	case 3:
		fmt.Println("The value is 3")
	default:
		fmt.Println("Default!!!")
	}

	switch a {
	case 1:
		fallthrough // This is a keyword for continue the evaluation
	case 2:
		fallthrough
	case 3:
		fallthrough
	case 4:
		fallthrough
	case 5:
		fmt.Println("You are during the week")
	case 6:
		fallthrough
	case 7:
		fmt.Println("You are during the weekend!!!")
	default:
		fmt.Println("Not is a valid day")
	}

	switch b := 6; {
	case b >= 0 && b <= 5:
		fmt.Println("You are during the week")
	case b >= 6 && b <= 7:
		fmt.Println("You are during the weekend!!!")
	default:
		fmt.Println("Not is a valid day")
	}
}
