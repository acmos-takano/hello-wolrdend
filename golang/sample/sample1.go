package sample

import (
	"fmt"
	"sample2"
)

func s1(x, y int) string {
	s1 := fmt.Sprintf("%d:%d", x, y)
	return s1
}

func s2(x, y int) int {
	return 10
}

func s3(x, y int) int {
	return sample2.s2_2()
}
