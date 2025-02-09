package main

import "fmt"

func main() {

	person := map[int]string{// key dan value
		1: "wahyu",
		2: "dino",
	}
	fmt.Println(person)
	fmt.Println(person[1])
	fmt.Println(person[2])
}