package main

func main() {
	e1 := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 13,
			},
			Right: &TreeNode{
				Val: 4,
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
	}
	println(hasPathSum(e1, 22))

	e2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	println(hasPathSum(e2, 5))
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if isLeaf(*root) {
		if root.Val == targetSum {
			return true
		}
	}
	return target(root.Left, targetSum, root.Val, false) || target(root.Right, targetSum, root.Val, false)
}

func target(node *TreeNode, targetSum int, accumulatedSum int, wasLeaf bool) bool {
	if node == nil {
		if wasLeaf {
			return targetSum == accumulatedSum
		} else {
			return false
		}

	} else {
		currentSum := accumulatedSum + node.Val
		return target(node.Left, targetSum, currentSum, isLeaf(*node)) || target(node.Right, targetSum, currentSum, isLeaf(*node))
	}
}

func isLeaf(node TreeNode) bool {
	return node.Left == nil && node.Right == nil
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
