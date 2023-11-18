package main

import (
	"fmt"
)

func main(){
	var num int
	fmt.Print("number = ")
	fmt.Scan(&num)

	if num == 0 {
		fmt.Println("choose a number larger than 0")
		return
	}

	if num == 1 {
		fmt.Println("it's prime")
		return
	}

	for i := 2; i < num; i++ {
		if num % i == 0 {
			fmt.Println("it's not a prime")
			return
		}
	}

	fmt.Println("it's a prime")
}