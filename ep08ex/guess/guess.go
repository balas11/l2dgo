package main

import "fmt"

func prompt(set [8]int) bool {
	fmt.Println("----------------------------------")
	fmt.Println("if it is in the following")
	fmt.Println(set)
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
	var set1 = [8]int{1, 3, 5, 7, 9, 11, 13, 15}
	var set2 = [8]int{2, 3, 6, 7, 10, 11, 14, 15}
	var set3 = [8]int{4, 5, 6, 7, 12, 13, 14, 15}
	var set4 = [8]int{8, 9, 10, 11, 12, 13, 14, 15}

	fmt.Println("Guess a number between 1 and 15 and press return to continue...")
	fmt.Scanln()
	guess := 0

	if prompt(set1) {
		guess += 1
	}
	if prompt(set2) {
		guess += 2
	}
	if prompt(set3) {
		guess += 4
	}
	if prompt(set4) {
		guess += 8
	}

	fmt.Println("You guessed : ", guess)
}
