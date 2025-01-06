package main

import (
	"fmt"
	"math/rand"
	"time"
)

 var array [5]int

 func init() {
 	fmt.Println("Executing init function")
 	array = [5]int{1, 2, 3, 4, 5}
 }
 func main() {
 	fmt.Println("Executing main function")
 	fmt.Println(array)
 }
