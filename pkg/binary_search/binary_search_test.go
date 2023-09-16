package binary_search

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	values := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}
	target := 23

	result := Run(values, target)
	expected := 5
	
	if result != expected {
		t.Errorf("result: '%v', expected: '%v'", result, expected)
	}
}
