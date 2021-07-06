package main

import "fmt"

func main() {
	// arr := [4]string{"geeks", "for", "geeks", "GFG"}
	// var my_slice1 = arr[1:2]
	// my_slice2 := arr[0:]
	// my_slice3 := arr[:2]
	// my_slice4 := arr[:]

	// fmt.Println("My Array: ", arr)
	// fmt.Println("My Array: ", my_slice1)
	// fmt.Println("My Array: ", my_slice2)
	// fmt.Println("My Array: ", my_slice3)
	// fmt.Println("My Array: ", my_slice4)
	a := make([]int, 5)
	b := len(a)
	c := cap(a)
	fmt.Println(a, b, c)

}
