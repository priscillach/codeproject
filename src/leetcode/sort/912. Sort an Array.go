package _sort

import "math/rand"

// https://leetcode.com/problems/sort-an-array/

// FastSort fast sort
func FastSort(nums []int) {
	FastSortByRange(nums, 0, len(nums)-1)
}

func FastSortByRange(nums []int, left, right int) {
	if left >= right {
		return
	}
	cur := FastSortCore(nums, left, right)
	FastSortByRange(nums, left, cur-1)
	FastSortByRange(nums, cur+1, right)
}

func FastSortCore(nums []int, left, right int) int {
	randIdx := left + rand.Intn(right-left+1)
	tmp := nums[randIdx]
	nums[randIdx] = nums[right]
	for left < right {
		for left < right && nums[left] <= tmp {
			left++
		}
		nums[right] = nums[left]
		for left < right && nums[right] >= tmp {
			right--
		}
		nums[left] = nums[right]
	}
	nums[right] = tmp
	return left
}

// BubbleSort bubble sort
func BubbleSort(nums []int) {
	for i := len(nums) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				nums[j+1], nums[j] = nums[j], nums[j+1]
			}
		}
	}
}

// HeapSort heap sort
func HeapSort(arr []int) {
	n := len(arr)

	// Build a max-heap
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// Extract elements from the heap one by one
	for i := n - 1; i >= 0; i-- {
		// Move the current root (maximum element) to the end
		arr[0], arr[i] = arr[i], arr[0]

		// Call heapify on the reduced heap
		heapify(arr, i, 0)
	}
}

func heapify(arr []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	// If the left child is larger than the root
	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	// If the right child is larger than the current largest
	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	// If the largest is not the root
	if largest != i {
		// Swap the largest element with the root
		arr[i], arr[largest] = arr[largest], arr[i]

		// Recursively heapify the affected sub-tree
		heapify(arr, n, largest)
	}
}

// MergeSort merge sort
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return mergeSortCore(left, right)
}

func mergeSortCore(left, right []int) []int {
	merged := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			merged = append(merged, left[i])
			i++
		} else {
			merged = append(merged, right[j])
			j++
		}
	}

	// Append any remaining elements
	merged = append(merged, left[i:]...)
	merged = append(merged, right[j:]...)

	return merged
}
