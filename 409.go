package main

import "fmt"

func longestPalindrome(s string) int {
	var containSingle bool
	temp := make(map[rune]int)
	for _, item := range s {
		temp[item] += 1
	}
	var result int
	for _, v := range temp {
		result += v / 2
		if v%2 != 0 {
			containSingle = true
		}
	}
	if containSingle {
		return 2*result + 1
	} else {
		return 2 * result
	}
}

func main() {
	fmt.Println(longestPalindrome("ababababa"))
}
