package main

import "fmt"

func main() {
	var p *int
	var x int = 11

	fmt.Printf("The value of p is %p\n", p)
	fmt.Printf("The address of x is %p\n\n\n", &x)

	p = &x
	fmt.Printf("The value of p is %p\n", p)
	fmt.Printf("The value stored at p is %d\n\n\n", *p)

	*p = 12
	fmt.Printf("The value of p is %p\n", p)
	fmt.Printf("The value stored at p is %d\n", *p)
	fmt.Printf("The value of x is %d\n\n\n", x)

	var ps *string = new(string)
	*ps = "Hello"
	fmt.Printf("The value stored at ps is %s\n", *ps)
	fmt.Printf("The value of ps is %p\n", ps)
}
