package algo

import "fmt"

func BinarySearchIterative(array []int, target int) int {
	left := 0
	right := len(array) - 1
	for left <= right {
		mid := int(float64(left + (right-left)/2.0))
		if array[mid] > target {
			right = mid - 1
			continue
		}
		if array[mid] < target {
			left = mid + 1
			continue
		}
		return mid
	}
	return -1
}

func BinarySearchRecursive(array []int, target int) int {
	left := 0
	right := len(array) - 1
	mid := int(float64(left + (right-left)/2.0))

	if left > right {
		return -1
	}

	if array[mid] == target {
		return mid
	}
	if array[mid] > target {
		right = mid - 1
	}
	if array[mid] < target {
		left = mid + 1
	}
	return BinarySearchRecursive(array[left:right+1], target)
}

func TestBinarySearch() {
	fmt.Println(BinarySearchIterative([]int{1, 2, 3, 4, 5, 6, 7, 8}, 4))
	fmt.Println(BinarySearchRecursive([]int{1, 2, 3, 4, 5, 6, 7, 8}, 4))
	fmt.Println(BinarySearchIterative([]int{1, 2, 2, 2, 5, 6, 7, 8}, 2))
	fmt.Println(BinarySearchRecursive([]int{1, 2, 2, 2, 5, 6, 7, 8}, 2))
	fmt.Println(BinarySearchIterative([]int{1}, 1))
	fmt.Println(BinarySearchRecursive([]int{1}, 1))
	fmt.Println(BinarySearchIterative([]int{}, 1))
	fmt.Println(BinarySearchRecursive([]int{}, 1))
}
