package main

import (
	"fmt"
	"time"
)

func main(){
	x := time.Now()
	fmt.Printf("current date = %d/%v/%v\n", x.Month(), x.Day(), x.Year())
	fmt.Printf("current date = %v", x.Format("1/2/2006"))
}