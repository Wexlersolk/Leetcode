package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := []rune(strs[0])
	for i := 1; i < len(strs); i++ {
		word := []rune(strs[i])
		for j := 0; j < len(prefix) && j < len(word); j++ {
			if prefix[j] != word[j] {
				prefix = prefix[:j]
				break
			}
		}
		if len(word) < len(prefix) {
			prefix = prefix[:len(word)]
		}
	}
	return string(prefix)
}

