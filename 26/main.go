package main

func removeDuplicates(nums []int) int {
	num := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == num {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		} else {
			num = nums[i]
		}

	}
	return len(nums)

}
