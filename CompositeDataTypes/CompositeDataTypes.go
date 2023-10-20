package main

import (
	"fmt"
)

func main() {
	//Arrays + Loops
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	//FOR ARRAY LOOP
	for i := 0; i < len(primes); i++ {
		fmt.Println("i", i)
		fmt.Println("val", primes[i])
	}
	//FOR LOOP
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("sum", sum)
	//WHILE LOOP
	n := 1
	for n < 5 {
		n *= 2
	}

	for {
		sum++ // repeated forever
		if sum > 13 {
			break //EXIT LOOP
		}
	}
	fmt.Println(sum) // never reached

	sum = 0
	for i := 1; i < 5; i++ {
		if i%2 != 0 { // skip odd numbers
			continue
		}
		sum += i
	}
	fmt.Println(sum) // 6 (2+4)

	//Slices
	var s []int = primes[1:4]
	fmt.Println(s)
	//Var Slices
	//Hash Tables
	//Maps
	//Structs

	fmt.Printf("Sprintf ")

}
