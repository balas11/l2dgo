package main

import "fmt"

func main() {
	//loop 1
	// for i := 0; i < 10; i++ {
	// 	fmt.Println("i = ", i)
	// }

	// loop 2
	// i := 0
	// for ; i < 10; i++ {
	// 	fmt.Println("i = ", i)
	// }

	// loop 3
	// i := 0
	// for i < 10 {
	// 	fmt.Println("i = ", i)
	// 	i++
	// }

	// loop 4
	i := 0
	for {
		fmt.Println("i = ", i)
		if i == 9 {
			break
		}
		i++
	}

}
