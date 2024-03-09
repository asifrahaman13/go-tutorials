package main

import (
	"fmt"
	"strings"
)

func total(length int, width int) int {
	return length + width
}

// Functions with variable number of inputs is also known as the variadic functions.

func joinstr(elements ...string) string {
	return strings.Join(elements, "-")
}

func main() {
	var x uint8 = 34
	fmt.Println(x)

	a := 34.5
	b := 45.4
	fmt.Println(a + b)

	const NAME = "My name" // const data are those data which are actually constant throughout the program.
	fmt.Println(NAME)

	var v int = 800

	if v < 1000 {
		fmt.Printf("v is less than 1000")
	} else {
		fmt.Print("V is greater thatn 1000")
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("Hello world\n")
	}

	fmt.Println(total(45, 66))

	fmt.Println(joinstr())
	fmt.Println(joinstr("G", "E", "E", "k", "S"))

	arr := [4]string{"Geek", "G", "E", "yellow"}
	fmt.Println(arr)

	my_array := [...]int{100, 200, 300, 400, 500}
	fmt.Println("Original array(Before):", my_array)

	myvar :="England"
	for index, s:=range myvar{
		fmt.Println(s, index)
	}
}
