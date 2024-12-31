package main

func romanToInt(s string) int {
	var sum int
	rti := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	rune := []rune(s)
	for i := len(rune) - 1; i >= 0; i-- {
		if i > 0 && rti[rune[i-1]] < rti[rune[i]] {
			sum += rti[rune[i]] - rti[rune[i-1]]
			i--
		} else {
			sum += rti[rune[i]]
		}
	}

	return sum
}
