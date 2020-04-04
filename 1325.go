package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return root
	}
	for {
		if isLeaf(root) && root.Val == target {
			return nil
		}
		deleted := new(bool)
		traverse(root, target, deleted)
		if !*deleted {
			break
		}
	}
	return root
}

func traverse(root *TreeNode, target int, deleted *bool) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		return
	}
	if root.Left != nil && isLeaf(root.Left) && root.Left.Val == target {
		*deleted = true
		root.Left = nil
	}
	if root.Right != nil && isLeaf(root.Right) && root.Right.Val == target {
		*deleted = true
		root.Right = nil
	}
	traverse(root.Left, target, deleted)
	traverse(root.Right, target, deleted)
	return
}

func isLeaf(node *TreeNode) bool {
	if node.Left == nil && node.Right == nil {
		return true
	}
	return false
}
