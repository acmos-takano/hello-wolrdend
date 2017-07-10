package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("test")
	fmt.Println("test2")

	time := time.Duration(100) * time.Second

	fmt.Println(time)
}
