package main

import (
	"fmt"
)

func main(){
	a := 10
	b := 20

	fmt.Printf("a = %v, b = %v\n", a, b)

	a,b = b,a

	fmt.Println("variables after swapping:")
	fmt.Printf("a = %v, b = %v", a, b)
}