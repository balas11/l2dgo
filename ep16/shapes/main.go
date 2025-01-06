package main

import (
	"fmt"
	"shapes/rect"
)

func main() {

	r := rect.New(10, 5)
	println(r.Area())
	println(r.Perimeter())
	println(r.IsSquare())
	r.Scale(2)
	fmt.Println(r.Length(), r.Width())
	fmt.Println(r)
}
