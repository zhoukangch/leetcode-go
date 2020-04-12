package main

import "container/list"

func minAddToMakeValid(S string) int {
	lis := list.New()
	count := 0
	for _, item := range S {
		if string(item) == ")" && lis.Len() == 0 {
			count++
		} else if string(item) == ")" && lis.Back().Value.(string) == "(" {
			lis.Remove(lis.Back())
		} else if string(item) == "(" {
			lis.PushBack(string(item))
		}
	}
	return count + lis.Len()
}
