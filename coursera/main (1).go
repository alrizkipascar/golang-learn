package main

import (
	"fmt"
	"math"
)

func main() {
	var floatingNum float64
	fmt.Println("Type a number")
	_, err := fmt.Scan(&floatingNum)

	if err != nil {
		fmt.Println("Something went wrong")
		return
	}

	fmt.Println(math.Trunc(floatingNum))
}
