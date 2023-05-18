package main

import "fmt"

func main() {
	var chicken map[string]int
	chicken = map[string]int{}
	chicken["januari"] = 50
	chicken["februari"] = 40                   //key
	fmt.Println("januari", chicken["januari"]) //januari
	fmt.Println("mei", chicken["mei"])         //mei 0

}
