/*
Function Multiple Return
Umumnya hanya memiliki satu buah nilai balik saja.
Jika ada kebutuhan di mana data yang dikembalikan harus banyak,
biasanya digunakanlah tipe seperti map, slice, atau struct sebgaia nilai balik.
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	area, circum := calculate(3.5)
	fmt.Println(area)
	fmt.Println(circum)
}

func calculate(d float64) (float64, float64) { // hitung luas
	var area = math.Pi * math.Pow(d/2, 2) // hitung keliling
	var circumference = math.Pi * d
	// kembalikan 2 nilai
	return area, circumference
}
