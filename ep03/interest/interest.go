package main

import "fmt"

// Step 1
//func main() {
//	var principle int
//	var interestRate float32
//	var years int
//	var interest float32
//
//	principle = 1000
//	interestRate = 6.5
//	years = 2
//
//	var name string = "பாலா"
//
//	fmt.Println(name)
//	interest = float32(principle) * (interestRate / 100) * float32(years)
//	fmt.Printf("Interest: %.2f\n", interest)
//}

//Step 2 - initialize variables while declaring
func main() {
	var principle, years int = 1000, 2

	var interestRate float32 = 6.5

	interest := float32(principle) * float32(years) * (interestRate / 100)

	fmt.Printf("The interest that you earned is %.2f\n", interest)
}
