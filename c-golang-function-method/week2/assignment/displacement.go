package main

import (
	"fmt"
)

func inputNumber() float64 {
	var number float64

	for {
		var err error
		if _, err = fmt.Scanln(&number); err == nil {
			break
		} else {
			fmt.Println("Error:", err)
		}
	}

	return number
}

func GenDisplaceFn(acceleration float64, initialVelocity float64, initialDisplacement float64) func(time float64) float64 {
	return func(time float64) float64 {
		return 0.5*acceleration*time*time + initialVelocity*time + initialDisplacement
	}
}

func main() {
	fmt.Print("Please enter Acceleration number: ")
	a := inputNumber()
	fmt.Print("Please enter Initial Velocity number: ")
	vo := inputNumber()
	fmt.Print("Please enter Initial Displacement number: ")
	so := inputNumber()
	fmt.Print("Please enter time number: ")
	time := inputNumber()

	computeDisplacement := GenDisplaceFn(a, vo, so)
	displacement := computeDisplacement(time)
	fmt.Println(displacement)
}
