package main

import "fmt"

func max[T ~int | string](a T, b T) T {
	if a > b {
		return a
	}
	return b
}

type ID int
type Number interface {
	int8 | int16 | int32 | int64 | float32 | float64 | int | string
}

func maxSlice[T Number](s []T) T {
	max := s[0]
	for _, v := range s {
		if v > max {
			max = v
		}
	}
	return max
}

type Kid struct {
	name string
	age  int
}

func find[T comparable](s []T, e T) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}

func maxAny[T any](s []T, bigger func(a, b T) bool) T {
	max := s[0]
	for _, v := range s {
		if bigger(v, max) {
			max = v
		}
	}
	return max
}

func main() {
	fmt.Println(max(1, 2))
	fmt.Println(max("a", "b"))

	// max("1", "a")

	fmt.Println(maxSlice([]int{1, 2, 3}))
	fmt.Println(maxSlice([]string{"a", "b", "c"}))
	fmt.Println(maxSlice([]float64{1.0, 2.0, 3.0}))

	var id1 = ID(1)
	var id2 = ID(2)
	fmt.Println(max(id1, id2))

	fmt.Println(find([]int{1, 2, 3}, 2))
	fmt.Println(find([]string{"a", "b", "c"}, "b"))

	kids := []Kid{
		{"a", 1},
		{"b", 2},
		{"c", 3},
	}

	fmt.Println(find(kids, Kid{"b", 2}))

	fmt.Println(maxAny([]int{1, 2, 3}, func(a, b int) bool {
		return a > b
	}))

	fmt.Println(maxAny(kids, func(a, b Kid) bool {
		return a.age > b.age
	}))

}
