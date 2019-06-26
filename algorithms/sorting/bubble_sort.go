package sorting

// BubbleSort sorts the slice
func BubbleSort(nums []int) []int {
	n := len(nums)

	// Run the outer loop n - 1 times
	// At the end of each iteration, we have bubbled up ith smallest item
	for i := 1; i <= n-1; i++ {

		// Run the inner loop on the rest of unsorted slice
		for j := n - 1; j >= i; j-- {

			// If value at index j is smaller than the one at j-1, swap them
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}

	return nums
}
