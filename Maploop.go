// Go program to illustrate how
// to iterate the map using for
// rang loop

package main

import "fmt"

// Main function
func main() {

	// Creating and initializing a map
	map1 := map[int]string{

		90: "Dog",
		91: "Cat",
		92: "Cow",
		93: "Bird",
		94: "Rabbit",
	}

	// Iterating map using for rang loop
	for id, pet := range map1 {

		fmt.Println(id, pet)
	}
}
