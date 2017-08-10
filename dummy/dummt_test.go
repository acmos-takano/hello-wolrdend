package dummy

import (
	"fmt"
	"testing"
)

func TestSaSamp_Aaa(t *testing.T) {
	fmt.Println("aaa")
	fmt.Println("bbb")
	fmt.Println("ccc")
	fmt.Println("ddd")

	t.Fatalf("aaa\n")
}
