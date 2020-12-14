package main

import "fmt"

func main() {
	node3 := &TreeNode{3, nil, nil}
	node2 := &TreeNode{2, nil, nil}
	tree1 := &TreeNode{1, node2, node3}
	tree2 := &TreeNode{1, node2, node3}

	// node2 := &TreeNode{2, nil, nil}
	// tree1 := &TreeNode{1, node2, nil}
	// tree2 := &TreeNode{1, nil, node2}

	fmt.Println(isSameTree(tree1, tree2))
}

// TreeNode represents a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	// Nil nodes do not have values, so nil cases must be handled before
	// comparing node values.

	// Case 1: Nil values, no match
	if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	}

	// Case 2: Nil values, match
	if p == nil && q == nil {
		return true
	}

	// Case 3: Non-nil values, no match
	if p.Val != q.Val {
		return false
	}

	if !isSameTree(p.Left, q.Left) {
		return false
	}
	if !isSameTree(p.Right, q.Right) {
		return false
	}

	// All matches
	return true
}
