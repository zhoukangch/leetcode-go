package main

func maxDistToClosest(seats []int) int {
	maxLen := -1
	for index, seat := range seats {
		if seat != 0 {
			continue
		}
		var offset int
		for offset = 1; index-offset >= 0 || index+offset < len(seats); offset++ {

			if index-offset >= 0 && seats[index-offset] != 0 {
				break
			}
			if index+offset < len(seats) && seats[index+offset] != 0 {
				break
			}
		}
		if offset > maxLen {
			maxLen = offset
		}
	}
	return maxLen
}
