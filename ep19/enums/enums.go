package main

import "fmt"

func Map(data []int, fn func(int) int) []int {
	mapped := make([]int, len(data))

	for i, v := range data {
		mapped[i] = fn(v)
	}
	return mapped
}

func main() {
	input := []int{1, 2, 3, 4, 5}

	output := Map(input, func(i int) int {
		return i * 2
	})

	fmt.Println(output)

	output = Map(input, func(i int) int {
		return i * i
	})

	fmt.Println(output)
}
