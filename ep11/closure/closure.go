package main

import "fmt"

// Store func in a variable
// func main() {

// 	greetInEnglish := func(name string) string {
// 		return "Hello, " + name
// 	}

// 	greetInSpanish := func(name string) string {
// 		return "Hola, " + name
// 	}

// 	fmt.Println(greetInEnglish("John"))
// 	fmt.Println(greetInSpanish("John"))

// }

// Pass func as parameter

// func calculate(operation func(int, int) int, arg1 int, arg2 int) int {
// 	return operation(arg1, arg2)
// }

type binaryOperation func(int, int) int

func calculate(operation binaryOperation, arg1 int, arg2 int) int {
	return operation(arg1, arg2)
}

func main() {
	adder := func(x, y int) int {
		return x + y
	}
	fmt.Println(calculate(adder, 2, 3))

	diff := calculate(func(x, y int) int {
		return x - y
	}, 2, 3)
	fmt.Println(diff)
}
