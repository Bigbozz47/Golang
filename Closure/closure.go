/*
	Function Closure
	==>adalah sebuah fungsi yang disimpan dalam variabel.

	Dengan menerapkan konsep tersebut,
	kita bisa membuat fungsi didalam fungsi atau bahkan membuat fungsi yang mengembalikkan fungsi.

*/

package main

import "fmt"

func main() {
	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}

	var numbers = []int{10, 3, 2, 2, 3, 1}
	var min, max = getMinMax(numbers)
	fmt.Println(min)
	fmt.Println(max)

}
