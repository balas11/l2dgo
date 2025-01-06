package main

import "fmt"

func sumton(n int) int {
	if n == 0 {
		return 0
	}
	return n + sumton(n-1)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func sumtonTail(n int) int {
	return sumtontr(0, n)
}

func sumtontr(acc int, n int) int {
	if n == 0 {
		return acc
	}
	return sumtontr(acc+n, n-1)
}

func factorialTail(n int) int {
	return facttr(1, n)
}

func facttr(acc int, n int) int {
	if n == 0 {
		return acc
	}
	return facttr(acc*n, n-1)
}

func gcd(x int, y int) int {
	if y == 0 {
		return x
	}
	return gcd(y, x%y)
}

func main() {
	fmt.Println(sumton(5))
	fmt.Println(factorial(5))
	fmt.Println(sumtonTail(5))
	fmt.Println(factorialTail(5))

	fmt.Println(gcd(26, 39))
}
