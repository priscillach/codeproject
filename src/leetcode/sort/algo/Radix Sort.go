package algo

import "leetcode/src/utils/mathhelper"

// RadixSortLSD
func RadixSortLSD(arr []int) []int {
	// Find the maximum number to determine the number of digits
	maxNum := mathhelper.Max(arr...)

	// Perform counting sort for each digit (from LSD to MSD)
	exp := 1
	for maxNum/exp > 0 {
		countingSort(arr, exp)
		exp *= 10 // Move to the next digit
	}
	return arr
}

// Counting sort for a specific digit (exp)
func countingSort(arr []int, exp int) {
	n := len(arr)
	output := make([]int, n) // Output array to store sorted elements
	count := make([]int, 10) // Count array to store frequency of digits (0-9)

	// Count the frequency of each digit at the current place (exp)
	for i := 0; i < n; i++ {
		index := (arr[i] / exp) % 10
		count[index]++
	}

	// Update count[i] to store the actual position of the digit in the output array
	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	// Build the output array by placing elements in their correct positions
	for i := n - 1; i >= 0; i-- {
		index := (arr[i] / exp) % 10
		output[count[index]-1] = arr[i]
		count[index]--
	}

	// Copy the sorted elements back to the original array
	for i := 0; i < n; i++ {
		arr[i] = output[i]
	}
}
