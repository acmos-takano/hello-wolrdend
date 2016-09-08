package sample

import (
	"fmt"

	"../sample2" // サンプル：自分の別パッケージを指定する場合の相対パス。ただしGOPATH配下では使えない
)

// S1:string
func S1(x, y int) string {
	s1 := fmt.Sprintf("%d:%d", x, y)
	return s1
}

// S2: int
func S2(x, y int) int {
	return 100
}

// S3: string
func S3(x, y int) string {
	s1 := fmt.Sprintf("%d:%d:%d", x, y, sample2.S2_2())
	return s1
}
