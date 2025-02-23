package main

func strStr(haystack string, needle string) int {
	nrune := []rune(needle)
	hrune := []rune(haystack)
	count := -1
	for i, letter := range hrune {
		if letter == nrune[0] {
			count = i
			for j := 1; j < len(nrune); j++ {
				if i+j < len(hrune) {
					if hrune[i+j] != nrune[j] {
						count = -1
					}
				} else {
					count = -1
				}
			}

			if count != -1 {
				return count
			}
		}

	}
	return count

}
