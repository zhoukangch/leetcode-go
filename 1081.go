package main

import (
	"container/list"
	"fmt"
)

func smallestSubsequence(text string) string {
	count := make([]int, 26, 26)
	for _, item := range text {
		count[item-'a']++
	}
	stack := list.New()
	dup := make([]bool, 26, 26)
	for _, item := range text {
		count[item-'a']--
		if dup[item-'a'] {
			continue
		}
		for stack.Len() > 0 {
			pre := stack.Back().Value.(int32)
			if pre <= item {
				break
			}
			if count[pre-'a'] <= 0 {
				break
			}
			dup[pre-'a'] = false
			stack.Remove(stack.Back())
		}
		stack.PushBack(item)
		dup[item-'a'] = true
	}
	var result []int32
	for stack.Len() > 0 {
		v := stack.Remove(stack.Front())
		result = append(result, v.(int32))
	}
	return string(result)
}

func main() {
	fmt.Println(smallestSubsequence("cdadabcc"))
}
