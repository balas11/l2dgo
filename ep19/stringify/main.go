package main

import (
	"fmt"
	"stringify/kid"
)

func main() {
	k := kid.New("John", 10)
	fmt.Println(k)
}
