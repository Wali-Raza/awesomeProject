package main

import "fmt"

func main() {
	var b bool
	//If no value is initilized then compiler willl consider it as false type
	var c bool = true
	var e bool = false
	fmt.Printf("%T  %v", b, b)
	fmt.Println()
	fmt.Printf("%T  %v", c, c)
	fmt.Println()
	fmt.Printf("%T  %v", e, e)
}
