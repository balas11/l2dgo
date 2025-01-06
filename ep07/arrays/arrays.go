package main

import "fmt"

func main() {

	var natural [10]int //filled with default int value 0

	fmt.Println(natural)

	natural[9] = 100
	fmt.Println(natural[9])

	for i := 0; i < 10; i++ {
		natural[i] = i + 1
	}
	fmt.Println(natural)

	sum := 0
	for i := 0; i < len(natural); i++ {
		sum += natural[i]
	}

	average := float32(sum) / float32(len(natural))
	fmt.Println(sum, average)

}
