package main

import "fmt"

func main() {
	var value1 = false && true
	var value2 = true || false

	fmt.Println(value1)
	fmt.Println(value2)
}
