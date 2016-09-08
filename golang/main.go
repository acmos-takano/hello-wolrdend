package main

import (
	"fmt"

	"./sample"
	"./sample2"
)

func main() {

	a := sample.S1(111, 999)
	b := sample2.S2_2()
	fmt.Println(a)
	fmt.Println(b)
}
