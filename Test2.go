package main

import "fmt"

var c int = 2 //Package level variable
func main() {
	b := 4
	//Function-level varibale
	fmt.Println("Package level Variable", c)
	fmt.Println("function level variable", b)
}
