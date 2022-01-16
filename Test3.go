package main

//type conversion
import "fmt"

func main() {
	a := 5
	fmt.Printf("%T,%v", a, a, "/n")
	var b float32 = float32(a)
	fmt.Println()
	fmt.Printf("%T  , %v", b, "/n", b)
}
