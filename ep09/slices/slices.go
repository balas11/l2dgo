package main

import "fmt"

func main() {

	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Original array : ", arr) //Original array :  [1 2 3 4 5]

	fmt.Println("------------------------------------")

	slice1 := arr[:]                                                     //slice points to the whole array
	fmt.Println("Slice1 (all) : ", slice1)                               //Slice1 (all) :  [1 2 3 4 5]
	fmt.Printf("Lenght : %d, Capacity : %d\n", len(slice1), cap(slice1)) //Lenght : 5, Capacity : 5

	fmt.Println("------------------------------------")

	slice2 := arr[:3]
	fmt.Println("Slice2 (0 to 2) : ", slice2)                            //Slice2 (0 to 2) :  [1 2 3]
	fmt.Printf("Lenght : %d, Capacity : %d\n", len(slice2), cap(slice2)) //Lenght : 3, Capacity : 5

	fmt.Println("------------------------------------")

	slice3 := arr[3:]
	fmt.Println("Slice3 (3 to end) : ", slice3)                          //Slice3 (3 to end) :  [4 5]
	fmt.Printf("Lenght : %d, Capacity : %d\n", len(slice3), cap(slice3)) //Lenght : 2, Capacity : 2

	fmt.Println("------------------------------------")

	slice4 := arr[1:4]
	fmt.Println("Slice 4 (1 to 3) : ", slice4)                           //Slice 4 (1 to 3) :  [2 3 4]
	fmt.Printf("Lenght : %d, Capacity : %d\n", len(slice4), cap(slice4)) //Lenght : 3, Capacity : 4

	fmt.Println("-------Grow with in capacity-----------")

	slice5 := append(slice4, 7)                                          //Grow with in capacity
	fmt.Println(slice5)                                                  //[2 3 4 7]
	fmt.Printf("Lenght : %d, Capacity : %d\n", len(slice5), cap(slice5)) //Lenght : 4, Capacity : 4
	fmt.Println(arr)                                                     //original array changed [1 2 3 4 7]

	fmt.Println("-------Grow Beyond capacity------------")

	slice6 := append(slice5, 8, 9)                                       //Grow beyond capacity
	fmt.Println(slice6)                                                  //[2 3 4 7 8 9]
	fmt.Printf("Lenght : %d, Capacity : %d\n", len(slice6), cap(slice6)) //Lenght : 6, Capacity : 8
	fmt.Println(arr)                                                     //[1 2 3 4 7] - original array is delinked from slice

	fmt.Println("-------Got a new array for the slice--------")
	slice6[0] = 99
	fmt.Println(slice6)                                                  //[99 3 4 7 8 9]
	fmt.Printf("Lenght : %d, Capacity : %d\n", len(slice6), cap(slice6)) //Lenght : 6, Capacity : 8
	fmt.Println(arr)

	fmt.Println("-------copy smaller target-------------") //[1 2 3 4 7]

	slice7 := []int{8, 9, 10}
	slice8 := make([]int, 2)
	copy(slice8, slice7)
	fmt.Println("slice7 : ", slice7) //slice7 :  [8 9 10]
	fmt.Println("slice8 : ", slice8) //slice8 :  [8 9]

	fmt.Println("-------copy larger target-------------")

	slice9 := make([]int, 5)
	copy(slice9, slice7)
	fmt.Println("slice7 : ", slice7)
	fmt.Println("slice9 : ", slice9)

}
