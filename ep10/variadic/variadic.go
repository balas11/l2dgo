package main

import "fmt"

//step 1
// func addem(sum *int, nums ...int) {
// 	*sum = 0
// 	for _, num := range nums {
// 		*sum += num
// 	}
// }

// func main() {
// 	var total int
// 	addem(&total, 1, 2, 3, 4, 5)
// 	fmt.Println(total)
// }

//step2

func addem(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	fmt.Println(addem(1, 2, 3, 4, 5))
}
