/*
283. Move Zeroes
https://leetcode.com/problems/move-zeroes/

Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.
*/

package movezeroes

// Time complexity: O(n)
// Space complexity: O(1)
func moveZeroes2(nums []int) {
	var k int
	for i := range nums {
		if nums[i] != 0 {
			if k != i {
				nums[k] = nums[i]
			}
			k++
		}
	}

	for i := k; i < len(nums); i++ {
		nums[i] = 0
	}
}
