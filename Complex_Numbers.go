package main

import "fmt"

func main() {
	var a complex64 = 1 + 2i
	var b complex64 = 1 + 2i

	fmt.Printf("Value => %v \nType=> %T \n", a, a)
	//fmt.Println(a + b)
	fmt.Println("Complex addition", a+b)
	fmt.Println("Complex Subtraction", a-b)
	fmt.Println("Complex Division", a/b)
	fmt.Println("Complex Multiplication", a*b)
}
