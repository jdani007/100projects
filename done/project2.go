package main

import (
	"fmt"
)

func main(){
	var num1, num2 int

	fmt.Print("What is num1? ")
	fmt.Scan(&num1)
	
	fmt.Print("What is num2? ")
	fmt.Scan(&num2)

	fmt.Println()
	fmt.Printf("Addition = %v\n", num1 + num2)
	fmt.Printf("Subtraction = %v\n", num1 - num2)
	fmt.Printf("Multiplication = %v\n", num1 * num2)
	fmt.Printf("Division = %v\n", num1 / num2)
} 