package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")
	languages := make(map[string]string)
	languages["js"] = "JavaScript"
	languages["python"] = "Python"
	languages["go"] = "golang"
	fmt.Println(languages)
	delete(languages, "js")
	fmt.Println(languages)

	// Loops are interesting in golang.
	for key, value:=range languages{
		fmt.Printf("For key %v, value is: %v\n", key, value)
	}
}
