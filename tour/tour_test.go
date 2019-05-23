package tour

import (
	"fmt"
	"testing"
)

func TestClosure(t *testing.T) {
	f := Fibonacci()
	x := f()
	y := f()
	z := f()
	if x != 0 || y != 1 || z != 1 {
		fmt.Errorf("Error with fiboacci closure, expected 0 1 1, got %v %v %v", x, y, z)
	}
}
