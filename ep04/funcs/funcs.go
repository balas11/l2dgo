package main

import "fmt"

func greet(name string) {
	message := "Hello " + name
	fmt.Println(message)
}

func calculateInterest(amount float32, years int, rate float32) float32 {
	return amount * float32(years) * rate / 100
}

func main() {
	greet("World")
	greet("Joe")

	interest := calculateInterest(1000.0, 5, 6.5)
	fmt.Printf("%.2f\n", interest)
}
