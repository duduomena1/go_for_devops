package math

import "testing"

func TestAdd(t *testing.T) {
	result := add(1, 2)
	if result != 3 {
		t.Errorf("add(1,2) = %d; want 3", result)
	}
}
