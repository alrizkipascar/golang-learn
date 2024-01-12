package main

import (
	"fmt"
	"math"

	. "OO/data"
)

type shape interface {
	DistToOrig() float64
	OffsetX() float64
}

type Point struct {
	x float64
	y float64
}

// Increase X coordiante
func (p *Point) OffsetX(v float64) {
	p.x = p.x + v
}

// Distance To Original
func (p Point) DistToOrig() float64 {
	t := math.Pow(p.x, 2) + math.Pow(p.y, 2)
	return math.Sqrt(t)
}

func main() {

	p1 := Point{x: 3, y: 4}

	fmt.Println("Before", p1)
	p1.OffsetX(5)
	fmt.Println("After", p1)

	fmt.Println("Distance", p1.DistToOrig())

	fmt.Println("Other Package", FuncSayHello())

	fmt.Println("Other Package X", X)

	fmt.Println("Other Package Secret", Secret)
	PrintX()
	// fmt.Println("Other Package PrintJ",)

	var p NewPoint
	p.InitMe(3, 4)
	p.Scale(2)
	p.PrintMe()
}
