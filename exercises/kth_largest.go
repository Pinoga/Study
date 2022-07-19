package exercises

import "sort"

func KthLargestOf(arr []int, k int) int {

	sort.Ints(arr)
	return arr[len(arr)-k]
}
