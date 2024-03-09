package main

import "fmt"


func greet() string{
	return "Hello world is the key"
}

func addition(a int,b int) int{
	return a+b
}

func main() {
  fmt.Println("Welcome to the function of golang")

  fmt.Println(greet())
  fmt.Println(addition(3,5))
}