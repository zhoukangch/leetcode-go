package main

import "fmt"

func numEnclaves(A [][]int) int {
	dp := make(map[int]bool)
	for row := range A {
		for col := range A[row] {
			if A[row][col] == 1 {
				dp[row*len(A)+col] = dfs(A, row, col, make(map[int]struct{}), dp)
			}
		}
	}
	var result int
	for _, v := range dp {
		if !v {
			result++
		}
	}
	return result
}

func dfs(A [][]int, row int, col int, dup map[int]struct{}, dp map[int]bool) bool {
	if row < 0 || row >= len(A) {
		return true
	}
	if col < 0 || col >= len(A[0]) {
		return true
	}
	if r, exist := dp[row*len(A)+col]; exist {
		return r
	}
	if A[row][col] == 0 {
		return false
	}
	dup[row*len(A)+col] = struct{}{}
	var r1, r2, r3, r4 bool
	if _, exist := dup[(row+1)*len(A)+col]; !exist {
		r1 = dfs(A, row+1, col, dup, dp)
		dup[(row+1)*len(A)+col] = struct{}{}
	}
	if _, exist := dup[(row-1)*len(A)+col]; !exist {
		r2 = dfs(A, row-1, col, dup, dp)
		dup[(row-1)*len(A)+col] = struct{}{}
	}
	if _, exist := dup[row*len(A)+col+1]; !exist {
		r3 = dfs(A, row, col+1, dup, dp)
		dup[row*len(A)+col+1] = struct{}{}
	}
	if _, exist := dup[row*len(A)+col-1]; !exist {
		r4 = dfs(A, row, col-1, dup, dp)
		dup[row*len(A)+col-1] = struct{}{}
	}
	if r1 || r2 || r3 || r4 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(numEnclaves([][]int{{0, 1, 1, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 0}}))
}

/*
0 0 0 0
1 0 1 0
0 1 1 0
0 0 0 0

0 1 1 0
0 0 1 0
0 0 1 0
0 0 0 0
*/
