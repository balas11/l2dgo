package main

import "fmt"

func main() {
	name := "John"

	switch len(name) {
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 3:
		fmt.Println(name, " is short length")
	case 4:
		fallthrough
	case 5:
		fallthrough
	case 6:
		fmt.Println(name, " is medium length")
	default:
		fmt.Println(name, " is long length")
	}
}
