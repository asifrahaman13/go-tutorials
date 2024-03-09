package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to the user input"

	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	println(reader)

	// comma ok syntax or the error ok syntax. In go there is no try catch blocks.
	input, _ := reader.ReadString('\n')
	println(input)
	println("The data type of the variable is %T",input)
}
