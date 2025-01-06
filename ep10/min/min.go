package main

import "fmt"

// step 1
// func min(a int, b int) {
// 	if a < b {
// 		fmt.Println("min : ", a)
// 	} else {
// 		fmt.Println("min : ", b)
// 	}
// }

// step 2
func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// min(10, 100)
	m := min(10, 100)
	fmt.Println("min : ", m)
}
