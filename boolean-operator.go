package main

import "fmt"

func main() {
	var nilaiA = 100
	var nilaiB = 80

	var result1 bool = nilaiA > 80
	var result2 bool = nilaiB >= 90

	var akhir bool = result1 && result2
	fmt.Println(akhir)


	var angka = 10
	var resulnya bool = angka >20
	var akhirnya = ! resulnya 
	fmt.Println(akhirnya)

}