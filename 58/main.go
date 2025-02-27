package main

func lengthOfLastWord(s string) int {
	words := []rune(s)
	count := 0
	flag := 0
	for _, char := range words {
		if char == ' ' {
			flag = count
		} else {
			if flag != 0 {
				count = 0
			}
			flag = 0
			count++
		}
	}
	if flag == 0 {
		return count
	} else {
		return flag
	}
}
