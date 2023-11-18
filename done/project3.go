package main

import (
	"fmt"
	"math"
)

func main(){
	var number float64
	fmt.Print("number = ")
	fmt.Scan(&number)

	fmt.Println("sqaure root =", int(math.Sqrt(number)))
}