package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func factorialTail(n int) int {
	return factorialtc(1, n)
}

func factorialtc(acc int, n int) int {
	if n == 0 {
		return acc
	}
	return factorialtc(acc*n, n-1)
}

func gcd(x int, y int) int {
	if y == 0 {
		return x
	}
	return gcd(y, x%y)
}

func main() {
	fmt.Println(factorial(5))
	fmt.Println(factorialTail(5))
	fmt.Println(gcd(10, 15))
	fmt.Println(gcd(128, 96))
}
