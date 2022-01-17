package main

import "fmt"

func main() {
	a := 4     //2^2
	a = a << 3 //2^2 * 2^3=32
	fmt.Println("Left shift  ", a)
	b := 8
	b = b >> 2
	fmt.Println("Right shift ", b)
	//c := 8
	//c &^ a
	var c int
	fmt.Println("bitwise", c&^a)
}
