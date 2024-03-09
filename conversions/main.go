package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to the pizza application.")
	fmt.Println("Please rate our pizza between 1 and 5.")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for reading", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added a one to the rating", numRating+1)
	}

	println("The numRating is %d", numRating)
	println("")
	fmt.Println("thanks for rarting", input)

}
