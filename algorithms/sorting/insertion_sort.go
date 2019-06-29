package sorting

// InsertionSort sorts the slice
func InsertionSort(nums []int) []int {
	n := len(nums)

	// Run the outer loop n - 1 times, from index 1 to n-1, as first element is already sorted
	// At the end of ith iteration, we have sorted list [0, i]
	for i := 1; i <= n-1; i++ {

		// Pick ith element and keep swapping with i-1th element if nums[i] < nums[i-1]
		j := i
		for j > 0 {

			// If value at index j is smaller than the one at j-1, swap them
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
			j -= 1
		}
	}

	return nums
}
