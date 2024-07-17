package slicesAndArrays

import (
	"reflect"
	"testing"
)

func sumVerify(t *testing.T, result, expected []int) {
	t.Helper()
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("SumAll was incorrect, got: %d, want: %v.", result, expected)
	}
}

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	result := Sum(numbers)
	expected := 15

	if result != expected {
		t.Errorf("Sum was incorrect, got: %d, want: %v.", result, expected)

	}
}
func TestSumAll(t *testing.T) {
	result := SumAll([]int{1, 2}, []int{0, 9})
	expected := []int{3, 9}
	sumVerify(t, result, expected)
}

func TestSumEverythingElse(t *testing.T) {
	/*	sumVerify := func(t *testing.T, result, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("SumAll was incorrect, got: %d, want: %v.", result, expected)
		}
	}*/
	t.Run("faz soma de alguns slices", func(t *testing.T) {
		result := SumEvereythingElse([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}
		sumVerify(t, result, expected)
	})

	t.Run("soma slices vazios de forma segura", func(t *testing.T) {
		result := SumEvereythingElse([]int{}, []int{3, 4, 5})
		expected := []int{0, 9}
		sumVerify(t, result, expected)
	})
}
