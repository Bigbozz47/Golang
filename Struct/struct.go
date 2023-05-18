package main

import "fmt"

type student struct {
	name  string
	grade int
}

func main() {
	var s1 student
	s1.name = "Delano Yusuf Habibie"
	s1.grade = 2
	fmt.Println("nama:", s1.name)   //januari
	fmt.Println("grade:", s1.grade) //mei 0

}
