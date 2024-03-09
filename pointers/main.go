package main

import "fmt"

type Employee struct {
	emp_id int
	name   string
}

func change(num *int) {
	*num = 11
}

func updateArray(arr *[5]int) {
	(*arr)[4] = 54 // Even the array will be changed in the main function since  here we are passing the address.
}

func main() {
	fmt.Println("Welcome to the class of pointers.")
	myNum := 23
	var ptr = &myNum // What is the address of the variable.
	fmt.Println("The value of actual pointer is: ", ptr)
	fmt.Println("The value of actual pointer is ", *ptr) // What is the actual value the pointer is refering to.

	// Double pointers
	var v int = 100
	var ptr1 *int = &v     // This means it is a pointer of integer data type.
	var ptr2 **int = &ptr1 // double ** means it is pointer to pointer.
	fmt.Println("The ptr1 is: ", ptr1)
	fmt.Println("The ptr2 is: ", ptr2)

	var a int = 10
	fmt.Println(a)
	var pa *int = &a
	change(pa) // The value of a changed to 11 since the pointer address to the address over the variable only.
	fmt.Println(a)

	arr := [5]int{78, 89, 45, 56, 14}
	updateArray(&arr)
	fmt.Println(arr)

	rohan := Employee{1, "Rahul"}
	fmt.Println(rohan.name)

	// First point to the address.
	pts := &rohan

	// We can also print using the pointer.
	fmt.Println((*pts).name)

	var ptr3 [8]*float64 
	fmt.Println(cap(ptr3))
	fmt.Println(len(ptr3))
}
