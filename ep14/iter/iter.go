package main

import "fmt"

type IntSlice []int

func (s IntSlice) Iter(doIt func(int)) {
	for _, v := range s {
		doIt(v)
	}
}

func (s IntSlice) Map(mapFunc func(int) int) {
	for i := 0; i < len(s); i++ {
		s[i] = mapFunc(s[i])
	}
}

func main() {
	s := IntSlice{1, 2, 3, 4, 5}
	sum := 0
	s.Iter(func(v int) {
		sum += v
	})
	fmt.Println(sum)
	
	s.Map(func(v int) int {
		return v * v
	})
	fmt.Println(s)

}
