package main

func main() {
	//e1 := []int{1}
	//printResult(sortedArrayToBST(e1))

	e2 := []int{-10, -3, -2, -1, 0, 1, 2, 5, 9, 10}
	printResult(sortedArrayToBST(e2))
}

func printResult(node *TreeNode) {
	if node != nil {
		println(node.Val)
		printResult(node.Left)
		printResult(node.Right)
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	center := len(nums) / 2
	if len(nums) == 1 {
		return &TreeNode{nums[0], nil, nil}
	}
	if len(nums) == 0 {
		return nil
	}
	leftArray := nums[:center]
	rightArray := nums[center+1:]
	left := sortedArrayToBST(leftArray)
	right := sortedArrayToBST(rightArray)
	node := &TreeNode{Val: nums[center],
		Left:  left,
		Right: right}

	return node
}
