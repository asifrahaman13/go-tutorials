package main

import "fmt"

func main() {
	defer fmt.Println("Hello to defer.")
	defer fmt.Println("Nonce")
	fmt.Println("Hello")
}