package iterators

import "testing"

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func TestRepeat(t *testing.T) {
	repetitions := Repeat("a", 5)
	expected := "aaaaa"
	if repetitions != expected {
		t.Errorf("got %v, want %v", repetitions, expected)
	}
}
