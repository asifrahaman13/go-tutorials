package main

import (
	"fmt"
	"time"
)

func display(str string) {
	for w := 0; w < 6; w++ {
		time.Sleep(1 * time.Second)
		fmt.Println(str)
	}
}

func init() {
	fmt.Println("Welcome to concurrency.")
}

func main() {
	fmt.Println("Hello world.")

	go display("Hello world with goroutine.")
	display("Hello world without goroutine.")
}
