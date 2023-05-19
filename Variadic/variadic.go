/*
	Function Variadic
	Go mengadopsi konsep variadic function atau pembuatan fungsi dengan paramater sejenis yang tak terbatas.
	Maksud tak terbatas adalah jumlah parameter yang disisipkan ketika pemanggilan fungsi bisa berapa saja.

*/

package main

import "fmt"

func main() {
	var avg = calculate(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
	var msg = fmt.Sprintf("Rata-rata : %.2f", avg)
	fmt.Println(msg)
}

func calculate(numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}
	var avg = float64(total) / float64(len(numbers)) // length
	return avg
}
