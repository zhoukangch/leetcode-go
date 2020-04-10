package main

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{
			Val: nums[0],
		}
	} else {
		mid := (len(nums) - 1) / 2
		root := &TreeNode{
			Val: nums[mid],
		}
		root.Left = sortedArrayToBST(nums[0:int(math.Max(float64(0), float64(mid)))])
		root.Right = sortedArrayToBST(nums[mid+1:])
		return root
	}
}
