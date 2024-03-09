package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome the the world")
	parseTime := time.Now()
	fmt.Println(parseTime)
	fmt.Println("I want to format", parseTime.Format(""))

	createDate := time.Date(2020, time.August, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(createDate)
}
