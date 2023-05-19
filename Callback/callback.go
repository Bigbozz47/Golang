/*
Function Callback
Fungsi bisa dijadikan sebagai tipe data variabel.
Dan sangat memungkinkan untuk menjadikannya sebagai parameter juga.
*/
package main

import "fmt"

func main() {
	var hasil = filter([]string{"ini", "data"}, func(each string) bool {
		return true
	})
	fmt.Println(hasil)
}

func filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}
