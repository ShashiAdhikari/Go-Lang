package main

import "fmt"

func main() {

	var map_1 map[int]int

	if map_1 == nil {
		fmt.Println("true")
	} else {
		fmt.Println("False")
	}

	map_2 := map[int]string{
		90: "DOG",
		91: "DB",
		92: "sds",
		94: "rabbit",
	}
	fmt.Println("Map-2: ", map_2)
}
