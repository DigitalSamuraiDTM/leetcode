package main

func main() {
	p1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	q1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	p2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 1,
		},
	}
	q2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 2,
		},
	}

	p3 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}

	q3 := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
		},
	}

	println(isSameTree(p1, q1)) // true
	println(isSameTree(p2, q2)) // false
	println(isSameTree(p3, q3)) // false
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q != nil ||
		p != nil && q == nil {
		return false
	}

	if p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}
		// is left nodes not same exist
		if p.Left == nil && q.Left != nil ||
			p.Left != nil && q.Left == nil {
			return false
		}
		// is right nodes not same exist
		if p.Right == nil && q.Right != nil ||
			p.Right != nil && q.Right == nil {
			return false
		}

		isLeftSubtreeSame := isSameTree(p.Left, q.Left)
		if isLeftSubtreeSame == false {
			return false
		}
		isRightSubtreeSame := isSameTree(p.Right, q.Right)
		if isRightSubtreeSame == false {
			return false
		}
	}
	return true
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
