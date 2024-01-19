package main

import (
	"fmt"
)

type Speaker interface {
	Speak()
}

type Speaker2 interface {
	Speaking()
}

type Dog struct{ name string }

type Dog2 struct{ name string }

func (d Dog) Speak() {

	fmt.Println(d.name)
}
func (d *Dog2) Speak2() {
	d.name = "new name"
	fmt.Println(d.name)
}

func (d *Dog) Speaking() {
	if d == nil {
		fmt.Println("<noise>")
	} else {
		fmt.Println(d.name)
	}
}

func (d *Dog2) Speaking2() {
	if d == nil {
		fmt.Println("<noise>")
	} else {
		fmt.Println(d.name)
	}
}

func main() {
	var s1 Speaker
	var d1 = Dog{"Brian"}
	s1 = d1
	s1.Speak()

	var s2 Speaker2
	var d2 *Dog
	s2 = d2
	s2.Speaking()

	// var s3 Speaker3
	// var name string = "test"
	var d3 = Dog2{"name"}

	d3.Speaking2()
	d3.Speak2()

}
