package main

func plusOne(digits []int) []int {
	return rec(digits, len(digits)-1)
}

func rec(digits []int, i int) []int {
	if digits[i] != 9 {
		digits[i] = digits[i] + 1
		return digits
	} else {
		digits[i] = 0
		if i == 0 {
			digits = append([]int{1}, digits...)
		} else {
			return rec(digits, i-1)
		}

	}
	return digits
}
