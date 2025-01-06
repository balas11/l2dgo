package main

import (
	"ecom/catalog"
	"fmt"
)

func main() {
	item1 := catalog.NewProduct("ABC Washing Soap", 10, 7.5)

	fmt.Printf("%s: ₹%.2f\n", item1.Name, item1.SellingPrice())

	fmt.Printf("Profit: ₹%.2f\n", item1.Profit())
}
