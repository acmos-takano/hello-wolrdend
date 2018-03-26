package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	fmt.Println("test")
	fmt.Println("test2")

	time := time.Duration(100) * time.Second

	fmt.Println(time)
	os.Exit(0)
}
