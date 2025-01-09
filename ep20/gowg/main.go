package main

import (
	"fmt"
	"sync"
	"time"
)

// var wg sync.WaitGroup

// func f(n int) {
//  defer wg.Done()
// 	for i := 0; i < n; i++ {
// 		fmt.Println("f:", i)
// 		time.Sleep(1 * time.Second)
// 	}
// }

// func g(n int) {
// 	defer wg.Done()
// 	for i := 0; i < n; i++ {
// 		fmt.Println("g:", i)
// 		time.Sleep(1 * time.Second)
// 	}
// }
// func main() {
// 	wg.Add(2)
// 	go f(10)
// 	go g(10)
// 	wg.Wait()
// }

func f(n int) {
	for i := 0; i < n; i++ {
		fmt.Println("f:", i)
		time.Sleep(1 * time.Second)
	}
}

func g(n int) {
	for i := 0; i < n; i++ {
		fmt.Println("g:", i)
		time.Sleep(1 * time.Second)
	}
}
func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		f(10)
	}()

	go func() {
		defer wg.Done()
		g(10)
	}()

	wg.Wait()
}
