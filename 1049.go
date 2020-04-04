package main

import (
	"fmt"
	"math"
)

func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, stone := range stones {
		sum += stone
	}
	maxCapacity := sum / 2
	dp := make([]int, maxCapacity+1, maxCapacity+1)
	for _, stone := range stones {
		for j := maxCapacity; j >= stone; j-- {
			dp[j] = int(math.Max(float64(dp[j]), float64(dp[j-stone]+stone)))
		}
	}
	return sum - 2*dp[maxCapacity]
}

func main() {
	s := []int{2, 7, 4, 1, 8, 1}
	fmt.Println(lastStoneWeightII(s))
}
