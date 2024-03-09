package main

import "fmt"

func main() {
	fmt.Println("If else in golang")

	loginCount := 10

	var result string

	// The blocks are important and needs to be in the same formatting as given.
	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "Something else"
	} else {
		result = "Result is exactly same."
	}

	fmt.Println(result)

	if num := 3; num < 10 {
		fmt.Println("The num is less than 10")
	} else {
		fmt.Println("Num is not less than 10")
	}

	diceNumber := 2
	// Switch caSe statement.
	switch diceNumber {
	case 1:
		fmt.Println("The dice value is 1")
	case 2:
		fmt.Println("The dice value is 2")
		fallthrough
	case 3:
		fmt.Println("The dice value is 3 through fallthrough")
	}
}
