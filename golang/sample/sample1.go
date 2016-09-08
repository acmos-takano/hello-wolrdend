package sample

import (
	"fmt"

	"../sample2" // サンプル：自分の別パッケージを指定する場合の相対パス。ただしGOPATH配下では使えない
)

// S1
func S1(x, y int) string {
	s1 := fmt.Sprintf("%d:%d", x, y)
	return s1
}

// S2
func S2(x, y int) int {
	return 10
}

// S3
func S3(x, y int) string {
	s1 := fmt.Sprintf("%d:%d:%d", x, y, sample2.S2_2())
	return s1
}
