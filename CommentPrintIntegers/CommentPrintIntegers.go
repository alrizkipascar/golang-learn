package main

import (
	"fmt"
	"log"
	"regexp"
	"runtime"
	"time"

	"rsc.io/quote"
)

// func timer(name string) func() {
//     start := time.Now()
//     return func() {
//         fmt.Printf("%s took %v\n", name, time.Since(start))
//     }
// }
func SomeFunction(list *[]string) {
    defer TimeTrack(time.Now())
    // Do whatever you want.
}

func TimeTrack(start time.Time) {
    elapsed := time.Since(start)

    // Skip this function, and fetch the PC and file for its parent.
    pc, _, _, _ := runtime.Caller(1)

    // Retrieve a function object this functions parent.
    funcObj := runtime.FuncForPC(pc)

    // Regex to extract just the function name (and not the module path).
    runtimeFunc := regexp.MustCompile(`^.*\.(.*)$`)
    name := runtimeFunc.ReplaceAllString(funcObj.Name(), "$1")

    log.Println(fmt.Sprintf("%s took %s", name, elapsed))
}
func main() {
	// comment
	var z int32 = 0; //comment on code
	/*
	Block Comment
	blablabla
	blablabla
	*/



	type Celcius float64;
	type IDperson int;
    var x int32 = 2;
    var y int16 = 3;
    
    fmt.Println(x + int32(y))
    fmt.Println(quote.Go())
    x = int32(y)
	// var j int;
	j := "Yo" //this is illegal btw
	fmt.Print("print",x + z)
	fmt.Println("println",x + z)
	fmt.Printf("printf " + j)

	// %s format string
	fmt.Print("\n")
	fmt.Printf("Sprintf %s", j)
	
	fmt.Sprintf("Sprintf %s" + j)

	//=======================Math Operations========================
	//Arithmetic
	fmt.Println("Arithmetic +",x + int32(y))
	fmt.Println("Arithmetic -",x - int32(y))
	fmt.Println("Arithmetic *",x * int32(y))
	fmt.Println("Arithmetic /",x / int32(y))
	fmt.Println("Arithmetic %",x % int32(y))
	fmt.Println("Arithmetic <<",x << int32(y))
	fmt.Println("Arithmetic >>",x >> int32(y))
	defer TimeTrack(time.Now())
	//Comparison
	fmt.Println("Comparison ==",x == int32(y))
	fmt.Println("Comparison !=",x != int32(y))
	fmt.Println("Comparison >",x > int32(y))
	fmt.Println("Comparison <",x < int32(y))
	fmt.Println("Comparison >=",x >= int32(y))
	fmt.Println("Comparison <=",x <= int32(y))
	defer TimeTrack(time.Now())
	//Boolean
	if x < 4 && y<4 {
		fmt.Println("Boolean &&",true)
	} else {
		fmt.Println("Boolean &&",false)
	}
	if x < 4 || y>4 {
		fmt.Println("Boolean ||",true)
	}	else {
		fmt.Println("Boolean ||",false)
	}
	defer TimeTrack(time.Now())
	
}