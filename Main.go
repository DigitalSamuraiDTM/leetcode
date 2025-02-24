package main

func main() {
	// solved without testing
}

func postorderTraversal(root *TreeNode) []int {
	out := []int{}
	if root != nil {
		if root.Left != nil {
			rekykler(root.Left, &out)
		}
		if root.Right != nil {
			rekykler(root.Right, &out)
		}
		out = append(out, root.Val)
	}
	return out
}

func rekykler(node *TreeNode, out *[]int) {
	if node != nil {
		if node.Left != nil {
			rekykler(node.Left, out)
		}
		if node.Right != nil {
			rekykler(node.Right, out)
		}
		*out = append(*out, node.Val)
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
