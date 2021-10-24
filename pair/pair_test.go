package pair

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	p := New(1, "one")

	fmt.Println(p)
}
