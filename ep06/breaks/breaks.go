package main

//step 1
// func main() {
// 	n := 27
// 	r := n % 10
// 	switch r { // only optional expression
// 	case 1, 3, 5, 7, 9:
// 		println("odd")
// 	case 0, 2, 4, 6, 8:
// 		println("even")
// 	default:
// 		println("Wont execute")
// 	}
// }

// step2
// func main() {
// 	n := 28

// 	switch r := n % 10; r { // both statement and expression
// 	case 1, 3, 5, 7, 9:
// 		println("odd")
// 	case 0, 2, 4, 6, 8:
// 		println("even")
// 	default:
// 		println("Wont execute")
// 	}
// }

// step3
func main() {
	n := 28

	switch r := n % 10; { // only optional expression
	case r == 1, r == 3, r == 5, r == 7, r == 9:
		println("odd")
	case r == 0, r == 2, r == 4, r == 6, r == 8:
		println("even")
	default:
		println("Wont execute")
	}
}
