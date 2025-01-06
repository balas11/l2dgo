package main

import "fmt"

const (
	Red   = 1
	Green = 2
	Blue  = 3
)

const (
	Sunday  = iota
	Monday  = iota
	Tuesday = iota
)

func main() {
	fmt.Printf("%d %d %d\n", Red, Green, Blue)
	fmt.Printf("%d %d %d\n", Sunday, Monday, Tuesday)
}
