package main

import (
	"fmt"
	"math"
)

func showAnimals(animals [15]string, setNumber int) {
	for idx, animal := range animals {
		if (idx+1)&int(math.Pow(2, float64(setNumber))) > 0 {
			fmt.Println(animal)
		}
	}
}

func prompt(animals [15]string, setNumber int) bool {
	fmt.Println("----------------------------------")
	fmt.Println("if it is in the following")
	showAnimals(animals, setNumber)
	fmt.Print("Press Y/y or else press N/n and enter : ")
	var resp string
	fmt.Scanf("%s", &resp)

	if resp == "Y" || resp == "y" {
		return true
	} else {
		return false
	}
}

func main() {
	var animals = [15]string{
		"Cat", "Dog", "Elephant", "Tiger", "Lion",
		"Monkey", "Giraffe", "Zebra", "Bear", "Wolf",
		"Fox", "Rabbit", "Deer", "Horse", "Kangaroo",
	}
	for _, animal := range animals {
		fmt.Println(animal)
	}
	fmt.Print("Guess one of the animals above ! Press return to continue...")
	fmt.Scanln()
	guess := 0
	for i := 0; i < 4; i++ {
		if prompt(animals, i) {
			guess += int(math.Pow(2, float64(i)))
		}
	}
	if guess != 0 {
		fmt.Println("You guessed :", animals[guess-1])
	}
}
