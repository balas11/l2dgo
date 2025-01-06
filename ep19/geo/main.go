package main

import (
	"fmt"
	"geo/shapes"
	"sort"
)

func main() {
	rectangles := shapes.Rectangles{
		shapes.Rectangle{Length: 10, Width: 5},
		shapes.Rectangle{Length: 20, Width: 10},
		shapes.Rectangle{Length: 15, Width: 7},
		shapes.Rectangle{Length: 7, Width: 12},
	}

	sort.Sort(rectangles)

	for _, rectangle := range rectangles {
		fmt.Println(rectangle)
	}
}
