package main

import "fmt"

// TreeNode represents a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var output []int

func main() {
	fmt.Println(rightSideView(&TreeNode{1, nil, nil}))
}

func dfs(root *TreeNode, depth int) {
	if root == nil {
		return
	}

	if depth > len(output) {
		output = append(output, root.Val)
	}

	dfs(root.Right, depth+1)
	dfs(root.Left, depth+1)
}

func rightSideView(root *TreeNode) []int {
	output = []int{}
	dfs(root, 1)
	return output
}
