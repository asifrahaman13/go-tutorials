package main

import (
	// "bytes"
	"fmt"
	"sort"
)

type Address struct {
	Name      string
	Address   string
	Telephone int
}

func addAllElements(slice_element []int) int {
	sum := 0
	for i := 0; i < len(slice_element); i++ {
		sum += slice_element[i]
	}
	return sum
}

func main() {
	arr := [][]string{{"geeks", "holland", "tokyo"}, {"geeks", "holland", "tokyo"}}
	fmt.Println(arr)
	arr2 := [...]int{9, 7, 6}
	fmt.Println(arr2)

	myslice := make([]int, 4, 7)
	fmt.Println(myslice)
	fmt.Printf("The length of the slice is: %d", len(myslice))

	for index := 0; index < len(myslice); index++ {
		fmt.Println(index)
	}

	for _, item := range myslice {
		fmt.Println(item)
	}

	source := []int{1, 2, 4, 5, 7}
	destination := make([]int, len(source))
	copy(destination, source[:])
	fmt.Println(destination)

	fmt.Printf("The sum of the elements is : %d.\n", addAllElements(source))

	// fmt.Println(bytes.Equal(arr, source))

	sort.Ints(source)
	fmt.Println(source)

	rahul := Address{Name: "Rahul", Address: "New Delhi", Telephone: 1424}
	fmt.Println(rahul)



	
}
