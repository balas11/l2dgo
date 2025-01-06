package main

import "fmt"

func minmax(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

func minmax2(a, b int) (min, max int) {
	if a < b {
		min = a
		max = b
	} else {
		min = b
		max = a
	}
	return
}

func stats(s []int) (int, float32) {
	sum := 0

	for _, v := range s {
		sum += v
	}
	avg := float32(sum) / float32(len(s))
	return sum, avg
}

func sliceminmax(s []int) (int, int) {
	min, max := s[0], s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}

func main() {
	n, x := minmax(9, 6)
	fmt.Println("min : ", n, ", max : ", x)

	marks := []int{75, 76, 100, 74, 76}

	total, average := stats(marks)
	fmt.Printf("total : %d, average : %.2f\n", total, average)

	n2, x2 := minmax2(9, 6)
	fmt.Println("min : ", n2, ", max : ", x2)

	scores := []int{74, 76, 100, 73, 77}

	min, max := sliceminmax(scores)
	fmt.Printf("min : %d, max : %d\n", min, max)
}
