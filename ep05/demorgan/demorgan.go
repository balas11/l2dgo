package main

import "fmt"

func main() {
	n := 5

	if n >= 3 && n <= 10 {
		fmt.Println(n, " is between 3 and 10")
	}

	k := 11
	if !(k >= 3 && k <= 10) {
		fmt.Println(k, " is not between 3 and 10")
	}

	//apply deMorgans law
	m := 2
	if m < 3 || m > 10 {
		fmt.Println(m, " is not between 3 and 10")
	}
}
