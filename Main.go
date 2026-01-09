package main

func main() {
	e1 := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: nil,
		},
	}
	printResult(inorderTraversal(e1))

	e2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 6,
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 8,
				Left: &TreeNode{
					Val: 9,
				},
			},
		},
	}

	printResult(inorderTraversal(e2))
}

func printResult(result []int) {
	for _, value := range result {
		println(value)
	}
	println("  ")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var result []int
	recursive(root, &result)
	return result
}

func recursive(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}
	if node.Left != nil {
		recursive(node.Left, result)
	}
	*result = append(*result, node.Val)
	if node.Right != nil {
		recursive(node.Right, result)
	}
}
