package main

import "fmt"

func findLengthOfLCIS(nums []int) int {
	var count int
	var max int
	for idx, num := range nums {
		if idx == 0 || num > nums[idx-1] {
			count++
		} else {
			if count > max {
				max = count
			}
			count = 0
		}
	}
	if count > max {
		max = count
	}
	return max
}

func main() {
	fmt.Println(findLengthOfLCIS([]int{1, 3, 5, 7}))
}
