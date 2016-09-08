package main

import (
	"fmt"

	"./sample"
	"./sample2"
)

func main() {

	a := sample.S3(111, 999)
	b := sample2.S2_3()
	fmt.Println(a)
	fmt.Println(b)
}
