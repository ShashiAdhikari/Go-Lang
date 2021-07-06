package main

import "fmt"

func main() {
	var mp map[string]int = map[string]int{
		"a": 5,
		"B": 9,
		"c": 10,
	}
	mp := make(map[string]int)
	fmt.Println(mp)
}
