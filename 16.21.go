package main

func findSwapValues(array1 []int, array2 []int) []int {
	a1sum := 0
	a2sum := 0
	temp := make(map[int]struct{})
	for _, item := range array1 {
		a1sum += item
	}
	for _, item := range array2 {
		temp[item] = struct{}{}
		a2sum += item
	}
	if (a1sum-a2sum)%2 != 0 {
		return []int{}
	}
	for index := range array1 {
		array1[index] -= (a1sum - a2sum) / 2
	}
	for _, item := range array1 {
		if _, exist := temp[item]; exist {
			return []int{item + (a1sum-a2sum)/2, item}
		}
	}
	return []int{}
}
