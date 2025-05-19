package main

func main() {
	e1 := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
	}
	var out1 = preorderTraversal(e1)
	for i := 0; i < len(out1); i++ {
		print(out1[i])
		print(" ")
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var out []int
	recursively(root, &out)
	return out
}

func recursively(node *TreeNode, outArray *[]int) {
	if node == nil {
		return
	}
	*outArray = append(*outArray, node.Val)
	if node.Left != nil {
		recursively(node.Left, outArray)
	}
	if node.Right != nil {
		recursively(node.Right, outArray)
	}
}
