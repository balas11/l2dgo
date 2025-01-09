package main

import (
	"fmt"
	"sync"
)

type stock struct {
	mu       sync.Mutex
	products map[string]int
}

func (s *stock) reduce(name string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.products[name] -= 1
}

func main() {
	var wg sync.WaitGroup

	s := stock{
		products: map[string]int{"apple": 2000, "banana": 1000},
	}

	doSales := func(name string, q int) {
		for i := 0; i < q; i++ {
			s.reduce(name)
		}
		wg.Done()
	}

	wg.Add(4)

	go doSales("apple", 1000)
	go doSales("apple", 1000)
	go doSales("banana", 500)
	go doSales("banana", 500)

	wg.Wait()

	fmt.Println(s.products)

}
