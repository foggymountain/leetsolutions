package rotatearray

// Given an integer array nums, rotate the array to the right by k steps, where k is non negative
func rotate(nums []int, k int) {
	if k == 0 {
		return
	}

	if k > len(nums) {
		k = k % len(nums)
	}

	rotated := []int{}
	for i, _ := range nums {
		l := (len(nums) - k + i) % len(nums)
		rotated = append(rotated, nums[l])
	}
	copy(nums, rotated)

}
