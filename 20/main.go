package main

func isValid(s string) bool {
	open := []rune{'{', '(', '['}
	closed := []rune{'}', ')', ']'}
	rune := []rune(s)
	if len(rune) == 1 {
		return false
	}
	if len(rune)%2 != 0 {
		return false
	}
mn:
	for i := 0; i < len(rune); i++ {
		for c := range closed {
			if rune[i] == closed[c] {
				if i == 0 {
					return false
				}
				foundMatch := false
				for k := 0; k < i; k++ {
					for o := range open {
						if rune[k] == open[o] && o == c {
							foundMatch = true
							break
						}
					}
				}
				if !foundMatch {
					return false
				}
				continue mn
			}
		}
	jloop:
		for j := i; j < len(rune); j++ {
			for c := range closed {
				if rune[j] == closed[c] {
					if j == c {
						continue mn
					} else {
						for k := i; k < j; k++ {
							for o := range open {
								if rune[k] == open[o] {
									if c == o {
										continue jloop
									}
								}
							}
						}
						return false
					}

				}
			}

		}
	}
	return true

}
