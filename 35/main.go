package main

func searchInsert(nums []int, target int) int {
	for i, number := range nums {
		if number == target {
			return i
		}
		if number > target {
			return i
		}
	}
	return len(nums)
}
