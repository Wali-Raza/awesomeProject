package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := 5
	fmt.Printf("%T %v ", a, a)
	fmt.Println()
	var b float32 = float32(a)
	fmt.Println()
	fmt.Printf("%T   %v", b, b)
	var c string = strconv.Itoa(a)
	fmt.Println()
	fmt.Printf(c)
}
