package quick_sort

import (
	"testing"
)

func equal(obtained_values []int , expected_values []int) bool {
	for i := 0; i < len(obtained_values); i++ {
        if obtained_values[i] != expected_values[i] {
			return false;
		}
    }
		return true;
}

// func TestSwap(t *testing.T){
// 	values := []int {1, 2, 3, 4, 5}
// 	swap(values, 1, 4)
// 	expected := []int{1, 5, 3, 4, 2}
// 	if !equal(values, expected) {
// 		t.Errorf("resultado '%v', esperado '%v'", values, expected)
// 	}
// }

// func TestLomuto(t *testing.T) {
// 	values := []int{7, 8, 1, 2, 90, 4, 65, 32}
// 	left_position := 0
// 	right_position := len(values) - 1

// 	result := Lomuto(values, left_position, right_position)
// 	expected := []int {4, 1, 2, 7, 90, 8, 65, 32}
// 	if !equal(result, expected) {
// 		t.Errorf("resultado '%v', esperado '%v'", result, expected)
// 	}
// }

func TestLomuto(t *testing.T) {
	values := []int{7, 8, 1, 2, 90, 4, 65, 32}
	left_index := 0
	right_index := len(values) - 1

	pivot_index := Lomuto(values, left_index, right_index)
	expected := 3
	if pivot_index != expected {
		t.Errorf("result '%v', expected '%v'", pivot_index, expected)
	}
}

func TestQuickSort(t *testing.T) {
	values := []int{7, 8, 1, 2, 90, 4, 65, 32}
	QuickSort(values, 0, len(values) - 1)
	expected := []int{1, 2, 4, 7, 8, 32, 65, 90}
		if !equal(values, expected) {
		t.Errorf("resultado '%v', esperado '%v'", values, expected)
	}

}
