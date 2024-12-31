package main

import "strconv"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	slice := strconv.Itoa(x)
	for i := 0; i < len(slice)/2; i++ {
		if slice[i] == slice[len(slice)-(i+1)] {
			continue
		} else {
			return false
		}
	}
	return true
}
