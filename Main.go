package main

func main() {
	e1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
	}
	println(isSymmetric(e1))

	e2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val: 3,
			},
		},
	}
	println(isSymmetric(e2))
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var leftNode *TreeNode = root.Left
	var rightNode *TreeNode = root.Right
	return checkSubtree(leftNode, rightNode)
}

func checkSubtree(leftTree *TreeNode, rightTree *TreeNode) bool {
	if leftTree == nil && rightTree != nil ||
		leftTree != nil && rightTree == nil {
		return false
	}
	if leftTree == nil && rightTree == nil {
		return true
	}
	// both tree exist...
	if leftTree.Val == rightTree.Val {
		leftRightCheck := checkSubtree(leftTree.Left, rightTree.Right)
		if leftRightCheck {
			rightLeftCheck := checkSubtree(leftTree.Right, rightTree.Left)
			return rightLeftCheck
		} else {
			return false
		}
	} else {
		return false
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
