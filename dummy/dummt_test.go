package dummy

import (
	"fmt"
	"testing"
)

func TestAaa(t *testing.T) {
	fmt.Println("aaa")

	t.Error("aaa")
}
