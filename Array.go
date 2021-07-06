package main

import "fmt"

func main() {
	var a []int
	for i := 0; i <= 9; i++ {
		if i%2 == 0 {
			a = append(a, i)
		}

	}

	fmt.Println(a)

}
