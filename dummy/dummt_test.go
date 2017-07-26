package dummy

import (
	"fmt"
	"testing"
)

func TestAaa(t *testing.T) {
	fmt.Println("aaa")
	fmt.Println("bbb")
	fmt.Println("ccc")
	fmt.Println("ddd")

	t.Error("aaa")
}
