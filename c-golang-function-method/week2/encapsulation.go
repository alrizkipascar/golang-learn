package main

import (
	"fmt"
)

func main() {

	var count int = 10
	for i := 0; i < count; i++ {

		if i == 2 {
			fmt.Println("Count kek", i)
		}
		if i == 5 {
			fmt.Println("Count buzz", i)
		} else {
			fmt.Println("Count", i)

		}
	}
}
