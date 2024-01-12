package main

import (
	"fmt"
)

func PrintHello() {
	fmt.Println("Hello, world.")
}

// Function Parameters
func foo(x int, y int) {
	fmt.Println(x * y)
}

// Function With Return Value
func foox(j int) int {
	return j + 1
}

// Function With Multiple Return Value
func foo2(x int) (int, int) {
	return x, x + 1
}

// Call by Value (Passed arguments are copied to params)
func CBV(y int) {
	y = y + 1
}

// Call by Reference (Passed arguments are copied to params)
func CBR(y *int) {
	*y = *y + 1
}

// passing array arguments
func arrarg(x [3]int) int {
	return x[0]
}

// passing array pointers
func arrpoint(x *[3]int) {
	(*x)[0] = (*x)[0] + 1
}

// Passing slices
func arrSlices(sli []int) {
	sli[0] = sli[0] + 1
}

func main() {
	PrintHello()
	foo(2, 3)
	y := foox(1)
	fmt.Println(y)
	a, b := foo2(3)
	fmt.Println(a, b)

	// not effective (time) but data encapsulation
	val := 2
	CBV(val)
	fmt.Println("call by value", val)

	//  effective (time) but no data encapsulation
	val2 := 2
	CBR(&val2)
	fmt.Println("call by refference", val2)

	arr := [3]int{1, 2, 3}
	fmt.Println("Passing array argument", arrarg(arr))

	arrp := [3]int{1, 2, 3}
	arrpoint(&arrp)
	fmt.Print("Passing array pointer", arrp)

	sl := []int{1, 2, 3}
	arrSlices(sl)
	fmt.Println("Passing array slices", sl)
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
