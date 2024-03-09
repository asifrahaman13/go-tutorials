package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to filesystem")

	content:= "This needs to be in the file."
    
	// Create a file.
	file, err:=os.Create("./mylcogofile.txt")

	if err != nil {
		panic(err)
	}
     
	// Write the string into the file. 
	length, err:=io.WriteString(file, content)

	if err != nil {
		panic(err)
	}
    
	// Read the length of the file and close the same. 
	fmt.Println("The length is ", length)
	defer file.Close()
}