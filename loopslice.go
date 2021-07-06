package main

import "fmt"

func main() {
	myslice := []string{"this", "is", "the", "tutorial", "of", "go", "language"}

	for e := 0; e < len(myslice); e++ {
		fmt.Println(myslice[e])
	}
}
