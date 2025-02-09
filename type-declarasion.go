package main

import "fmt"

func main() {
	type NoKTP string
	var ktpEko NoKTP = "dino"

	fmt.Println(ktpEko)
}
