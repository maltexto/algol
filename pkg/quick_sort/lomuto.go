package quick_sort

func Lomuto(values []int, left int, right int) int {
	swap := func (st_index int, nd_index int) {
		temp := values[st_index]
		values[st_index] = values[nd_index]
		values[nd_index] = temp
	}

	pivot := values[left]
	i := left

	for j := left + 1; j <= right; j++ {
		if (values[j] <= pivot) {
			i+=1;
			swap(i, j);
		}
	}
	swap(i, left)
	return i
}

func QuickSort(values []int, left int, right int) {
	if left < right {
		index_pivot := Lomuto(values, left, right);
		QuickSort(values, left, index_pivot - 1);
		QuickSort(values, index_pivot + 1, right);	
	}
}