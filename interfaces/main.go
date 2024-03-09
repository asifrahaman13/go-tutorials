package main

import "fmt"

// So the interfaces are basically the unimplemented methods. Next the struct should define the instances and then we need to implement the methods.

// var interface_name InterfaceName
// interface_name = StructDataType
// Next call the methods. interface_name.FirstMethod()

type rectangle interface {
	Area() float64
	Peremeter() float64
	Sample() string
}

type dimensions struct {
	length  float64
	breadth float64
}

// Sample implements rectangle.
func (r dimensions) Sample() string {
	return "Not implemented. "
}

// Peremeter implements rectangle.
func (r dimensions) Peremeter() float64 {
	panic("unimplemented")
}

func (r dimensions) Area() float64 {
	return 2 * (r.length + r.breadth)
}

func main() {

	var r rectangle

	r = dimensions{12, 34}

	fmt.Println(r.Area())
	fmt.Println(r.Sample())

	fmt.Println("Hello world.")

	myValue := "Hello this is a value"

	for index, s := range myValue {
		fmt.Println(index, string(s))
	}

	for index := 1; index <= 10; index++ {
		fmt.Println(index)
	}

	// Iterate over the number slice using numbers.
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, num := range numbers {
		fmt.Println(num)
	}
}
