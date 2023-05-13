package main

import "fmt"

func main() {
	var point = 88
	switch point {
	case 8:
		fmt.Println("perfect")
	case 9:
		fmt.Println("good")
	default:
		fmt.Println("not bad")
	}
}
