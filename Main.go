package main

func main() {

	exampleNode := &TreeNode{
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

	exampleOnlyRight := &TreeNode{
		Val:  4,
		Left: nil,
		Right: &TreeNode{
			Val:  5,
			Left: nil,
			Right: &TreeNode{
				Val:  6,
				Left: nil,
				Right: &TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}

	println(maxDepth(exampleNode))
	println(maxDepth(exampleOnlyRight))
	println(maxDepth(nil))
}

func maxDepth(root *TreeNode) int {
	max := 1
	if root == nil {
		return 0
	}

	maxLeft := maxDepth(root.Left)
	maxRight := maxDepth(root.Right)
	if maxLeft > maxRight {
		max += maxLeft
	} else {
		max += maxRight
	}
	return max
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
