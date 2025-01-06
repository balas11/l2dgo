package main

import "fmt"

func main() {
	var n byte = 16

	fmt.Println(n << 1) //multiply by 2
	fmt.Println(n >> 1) //divide by 2

	fmt.Println(15 & 1)
	fmt.Println(15 & 0)

	fmt.Printf("%b\n", n)
	fmt.Printf("%b\n", ^n)

}
