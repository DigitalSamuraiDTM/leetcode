package main

func main() {
	// root = [3,9,20,null,null,15,7]
	var example1 = TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}
	println(minDepth(&example1))
	// root = [2,null,3,null,4,null,5,null,6]
	var example2 = TreeNode{
		Val:  2,
		Left: nil,
		Right: &TreeNode{
			Val:  3,
			Left: nil,
			Right: &TreeNode{
				Val:  9,
				Left: nil,
				Right: &TreeNode{
					Val:  15,
					Left: nil,
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
		},
	}
	println(minDepth(&example2))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return minHeightRecursion(root, 0)
}

func minHeightRecursion(node *TreeNode, minHeight int) int {
	newHeight := minHeight + 1

	if node.Left == nil && node.Right == nil {
		return newHeight
	} else {
		if node.Left == nil {
			return minHeightRecursion(node.Right, newHeight)
		} else if node.Right == nil {
			return minHeightRecursion(node.Left, newHeight)
		}
		leftHeight := minHeightRecursion(node.Left, newHeight)
		rightHeight := minHeightRecursion(node.Right, newHeight)
		if leftHeight > rightHeight {
			return rightHeight
		} else {
			return leftHeight
		}
	}
}
