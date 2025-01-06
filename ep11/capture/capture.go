package main

import "fmt"

// func main() {
// 	v := 5

// 	countdown := func() int {
// 		if v--; v <= 0 {
// 			v = 0
// 		}
// 		return v
// 	}

// 	for i := 0; i < 10; i++ {
// 		fmt.Println(countdown())
// 	}
// }

func countdown(v int) func() int {
	return func() int {
		r := v
		if v--; v <= 0 {
			v = 0
		}
		return r
	}
}

func makeMultipliers(n int) func() int {
	i := 0
	return func() int {
		i++
		return n * i
	}
}

func main() {
	down := countdown(5)

	for i := 0; i < 10; i++ {
		fmt.Println(down())
	}

	nextTriple := makeMultipliers(3)
	fmt.Println(nextTriple())
	fmt.Println(nextTriple())
	fmt.Println(nextTriple())
}
