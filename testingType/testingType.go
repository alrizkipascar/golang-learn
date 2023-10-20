package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	type Celcius float64;
	type IDperson int;
    var x int32 = 200;
    var y int16 = 300;
    
    fmt.Println(x + int32(y))
    fmt.Println(quote.Go())
    x = int32(y)
    fmt.Println(x + int32(y))
}