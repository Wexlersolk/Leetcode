package main

func climbStairs(n int) int {
	if n-2 == 0 {
		return 2
	}
	if n-1 == 0 {
		return 1
	}
	return climbStairs(n-1) + climbStairs(n-2)
}

func climbStairs(n int) int {
	memo := make(map[int]int)
	return helper(n, memo)
}

func helper(n int, memo map[int]int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if val, exists := memo[n]; exists {
		return val
	}
	memo[n] = helper(n-1, memo) + helper(n-2, memo)
	return memo[n]
}
