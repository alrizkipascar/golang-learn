package data

import "fmt"

// var X int = 1
// var Y int = 2

func FuncSayHello() string {
	return "Hi from package dir1"
}

var X int = 4444
var Secret string = "kaokwojtrarwa"
var j int = 1

func PrintX() {
	fmt.Println("print j", j)
}

type NewPoint struct {
	x float64
	y float64
}

func (p *NewPoint) InitMe(xn, yn float64) {
	p.x = xn
	p.y = yn
	fmt.Println("First input", p.x, p.y)

}

func (p *NewPoint) Scale(v float64) {
	p.x = p.x * v
	p.y = p.y * v
}
func (p *NewPoint) PrintMe() {
	fmt.Println("After Scale", p.x, p.y)
}
