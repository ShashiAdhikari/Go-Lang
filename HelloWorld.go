// package main : The package “main” tells the Go compiler that the package should compile as an executable program instead of a shared library.


package main

import "fmt"

// fmt: fmt stands for the Format package. 
//This package allows to format basic strings, values, or anything and print them or collect user input from the console,
//or write into a file using a writer or even print customized fancy error messages. This package is all about formatting input and output.

func main() {
  //The main() function is a special type of function and it is the entry point of the executable programs. 
  //It does not take any argument nor return anything.
	fmt.Println("HEllo world!")
}
