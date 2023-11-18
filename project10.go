package main

import (
	"fmt"
)

func main(){
	var num int
	fmt.Print("number = ")
	fmt.Scan(&num)

	var output string
	result := 1

	for i := num; i >= 1; i-- {
		result = result * i
		if i == 1 {
			output = output + fmt.Sprint(i)
		} else {
			output = output + fmt.Sprint(i) + "*"
		}
	}

	fmt.Printf("Factorial = %v = %v\n", output, result)
}