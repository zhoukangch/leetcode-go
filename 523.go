package main

func checkSubarraySum(nums []int, k int) bool {
	if k == 0 {
		count := 0
		for _, num := range nums {
			if num != 0 {
				count = 0
			} else {
				count++
			}
			if count >= 2 {
				return true
			}
		}
		return false
	}
	m := make(map[int]int)
	m[0] = -1
	sum := 0
	for idx, num := range nums {
		sum += num
		temp := sum % k
		if s, exist := m[temp]; exist {
			if idx-s > 1 {
				return true
			}
		} else {
			m[temp] = idx
		}
	}
	return false
}
