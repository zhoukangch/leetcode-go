package main

func reconstructMatrix(upper int, lower int, colsum []int) [][]int {
	array1 := make([]int, len(colsum), len(colsum))
	array2 := make([]int, len(colsum), len(colsum))
	var temp int
	for index, col := range colsum {
		if col == 2 {
			array1[index] = 1
			array2[index] = 1
			upper--
			lower--
		} else if col == 1 {
			temp++
		}
	}
	if upper < 0 || lower < 0 {
		return [][]int{}
	}
	if upper+lower != temp {
		return [][]int{}
	}
	for index, col := range colsum {
		if col == 1 {
			if upper > 0 {
				array1[index] = 1
				upper--
			} else {
				array2[index] = 1
				lower--
			}
		}
	}
	return [][]int{array1, array2}
}
