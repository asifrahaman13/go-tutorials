package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}

	// Loop statement.
	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	for index, day := range days {
		fmt.Printf("Index is %v and value is %v", index, day)
	}

	roughValue := 3

	for roughValue < 10 {

		// Break the statement on certain conditions.
		if roughValue == 5 {
			break
		}
		fmt.Println("The value is: ", roughValue)
		roughValue++
	}
}
