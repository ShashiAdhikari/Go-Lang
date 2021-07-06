package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "Welcome to the online portal of GeeksforGeeks"
	res1 := strings.Count(str, "o")
	fmt.Println("\nResult 1: ", res1)
}
