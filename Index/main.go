package main

import "fmt"

// declaration operator. This is not allowed outside the functions and methods.
// jwtToken:= 3489

const LoginInto string = "thisisatypecastingstring" // First letter as capital means it is a public variable.

func main() {
	fmt.Println("Hello, World!")
	var username string = "I am the user"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var isSmall uint8 = 45
	fmt.Printf("Variable is of type: %T \n", isSmall)

	var anotherVariablke int // There is not garbage collector in golang. Int will by default point to 0 instead.
	fmt.Println(anotherVariablke)

	numberOfUsers := 45435
	fmt.Println(numberOfUsers)

	fmt.Println(LoginInto)
}
