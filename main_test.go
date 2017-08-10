package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	fmt.Println("aaa")
	fmt.Println("bbb")
	fmt.Println("ccc")
	fmt.Println("eee")
	fmt.Println("019")
	fmt.Println("dr-test14")

	t.Error("aaa")

}

func TestMain2(t *testing.T) {
	fmt.Println("aaa")
	fmt.Println("bbb")
	fmt.Println("ccc")
	fmt.Println("eee")
	fmt.Println("019")
	fmt.Println("dr-test14")

	// t.Error("aaa")

}
