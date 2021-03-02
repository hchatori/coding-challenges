package main

import "fmt"

func main() {
	// Example 1: 3
	// node5 := TreeNode{Val: 7, Left: nil, Right: nil}
	// node4 := TreeNode{Val: 15, Left: nil, Right: nil}
	// node3 := TreeNode{Val: 20, Left: &node4, Right: &node5}
	// node2 := TreeNode{Val: 9, Left: nil, Right: nil}
	// root := TreeNode{Val: 3, Left: &node2, Right: &node3}

	// Example 2: 3
	node5 := TreeNode{Val: 5, Left: nil, Right: nil}
	node4 := TreeNode{Val: 4, Left: nil, Right: nil}
	node3 := TreeNode{Val: 3, Left: nil, Right: &node5}
	node2 := TreeNode{Val: 2, Left: &node4, Right: nil}
	root := TreeNode{Val: 1, Left: &node2, Right: &node3}

	fmt.Println(maxDepth(&root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	currDepth := 1
	maxDepth := 0

	if root == nil {
		return 0
	}

	maxDepth = searchTree(root, currDepth, maxDepth)

	return maxDepth
}

func searchTree(root *TreeNode, currDepth int, maxDepth int) int {

	if root.Left == nil && root.Right == nil {
		if currDepth > maxDepth {
			maxDepth = currDepth
		}

		return maxDepth
	}

	if root.Left != nil {
		currDepth++
		maxDepth = searchTree(root.Left, currDepth, maxDepth)
		currDepth--
	}

	if root.Right != nil {
		currDepth++
		maxDepth = searchTree(root.Right, currDepth, maxDepth)
		currDepth--
	}

	return maxDepth
}
