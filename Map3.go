// using make function

package main

import "fmt"

func main() {
	var My_map = make(map[float64]string)
	fmt.Println(My_map)

	My_map[1.3] = "Rohit"
	My_map[1.5] = "Sumit"
	fmt.Println(My_map)
}
