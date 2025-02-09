package main

import (
	"fmt"
)

func main() {
	var month = [...] string{
		 "januari",
		 "februari",
		 "maret",
		 "april",
		 "mei",
		 "juni",
		 "juli",
		 "agustus",
		 "september",
		 "oktober",
		 "november",
		 "desember",

	}
	
	var slice1 = month[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	slice1[0] = "dino"
	fmt.Println(slice1)

	var slice2 = month[10:]
	fmt.Println(slice2)

	slice3 :=append(slice2,"dino")
	fmt.Println(slice3)
	slice3[1] = "bukan desember"
	fmt.Println(slice3)

	newSlice := make([]string, 3, 5)
	newSlice[0] = "wahyu"
	newSlice[1] = "dino"
	newSlice[2] = "lowi"

	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	iniArrya := [...] int {1,2,3,4,5}
	iniSlice := []int {1,2,3,34,5,}
	fmt.Println(iniArrya)
	fmt.Println(iniSlice)


	
}