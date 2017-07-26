package dummy

import (
	"fmt"
	"testing"
)

func TestAaa(t *testing.T) {
	fmt.Println("aaa")
	fmt.Println("bbb")

	t.Error("aaa")
}
