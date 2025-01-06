package main

import "fmt"

func main() {

	var marks [5]int = [5]int{60, 75, 90, 80, 95}
	fmt.Println(marks)

	scores := [5]int{60, 75, 90, 80, 95}
	fmt.Println(scores)

	var names = [5]string{
		"John",
		"Jane",
		"Jim",
		"Jack",
		"Jill",
	}
	fmt.Println(names)

	for i, name := range names {
		fmt.Println(i, " : ", name)
	}

	total := 0
	for _, score := range scores {
		total += score
	}
	fmt.Println("Total : ", total)
}
