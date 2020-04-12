package main

import (
	"fmt"
	"sort"
	"strconv"
)

func lexicalOrder(n int) []int {
	var result []string
	for i := 1; i <= n; i++ {
		result = append(result, strconv.Itoa(i))
	}
	sort.Strings(result)
	var temp []int
	for _, item := range result {
		i, _ := strconv.Atoi(item)
		temp = append(temp, i)
	}
	return temp
}

func main() {
	fmt.Println(lexicalOrder(13))
}
