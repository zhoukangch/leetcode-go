package main

import (
	"sort"
)

func nextGreatestLetter(letters []byte, target byte) byte {
	letters = append(letters, target)
	sort.Slice(letters, func(i, j int) bool {
		if letters[i] <= letters[j] {
			return true
		}
		return false
	})
	findLatestIndex := func(a []byte, t byte) int {
		var index int
		for index = 0; index < len(a); index++ {
			if a[index] > t {
				break
			}
		}
		return index - 1
	}
	i := findLatestIndex(letters, target)
	if i == len(letters)-1 {
		return letters[0]
	} else {
		return letters[i+1]
	}
}

func main() {
	nextGreatestLetter([]byte{'c', 'f', 'j'}, byte('a'))
}
