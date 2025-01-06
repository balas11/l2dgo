package main

import (
	"errors"
	"fmt"
)

//step 1 complicated
// func validate(name string, age int) error {
// 	if len(name) >= 3 && len(name) <= 15 {
// 		if age > 18 && age < 60 {
// 			return nil
// 		} else {
// 			return errors.New("Age should be between 18 and 60")
// 		}
// 	} else {
// 		return errors.New("Name should be between 3 and 15 characters")
// 	}
// }

// step 2
func validate(name string, age int) error {
	if len(name) < 3 || len(name) > 15 {
		return errors.New("Name should be between 3 and 15 characters")
	}
	if age < 18 || age > 60 {
		return errors.New("Age should be between 18 and 60")
	}
	return nil
}
func main() {
	err := validate("Jo", 12)
	if err != nil {
		fmt.Println("1. ", err)
	}
	err = validate("John", 12)
	if err != nil {
		fmt.Println("2. ", err)
	}

	err = validate("Jo", 20)
	if err != nil {
		fmt.Println("3. ", err)
	}

	err = validate("John", 20)
	if err != nil {
		fmt.Println("4. ", err)
	}
}
