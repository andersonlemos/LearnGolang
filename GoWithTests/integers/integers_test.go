package integers

import "testing"

func TestAdding(t *testing.T) {
	sum := Add(1, 2)
	expected := 4
	if sum != expected {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", sum, expected)
	}
}
