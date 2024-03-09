package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to video slices")
	var fruitList = []string{"apple", "tomato", "orange"}
	fmt.Println("Type of fruit list is: ", fruitList)
	fmt.Println(len(fruitList))
	fruitList = append(fruitList, "mango", "banana")
	fmt.Println(fruitList)


	highScores:=make([]int, 4)
	highScores[0]=54
	highScores[1]=24
	highScores[2]=27
	highScores[3]=29
	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)


	var index int=1;
	highScores=append(highScores[:index], highScores[index+1:]...)
	fmt.Println(highScores)
}
