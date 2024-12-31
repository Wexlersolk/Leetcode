package main

func main() {

}

func module(a int) int {
	if a < 0 {
		return -a
	}
	return a

}
func scoreOfString(s string) int {
	var result int
	for i := 0; i < len(s)-1; i++ {
		ascii1 := int(s[i])
		ascii2 := int(s[i+1])
		result += module(ascii1 - ascii2)
	}
	return result
}
