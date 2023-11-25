package main

import (
	"fmt"
)

func main() {
var character string
fmt.Print("character = ")
fmt.Scan(&character)

fmt.Printf("%s %d\n",character, character[0])

}