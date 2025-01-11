package main

func main() {
	e1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}
	e2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
		},
	}
	e3 := &ListNode{
		Val:  1,
		Next: nil,
	}

	out(deleteDuplicates(e1))
	out(deleteDuplicates(e2))
	out(deleteDuplicates(e3))
}

func out(head *ListNode) {
	if head == nil {
		println("NIL")
		return
	}
	println(head.Val)
	out(head.Next)
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	root := &ListNode{
		Val:  head.Val,
		Next: nil,
	}
	currentNode := root
	nextNode := head.Next

	for nextNode != nil {
		if nextNode.Val != currentNode.Val {
			currentNode.Next = &ListNode{
				Val:  nextNode.Val,
				Next: nil,
			}
			currentNode = currentNode.Next
		}
		nextNode = nextNode.Next
	}
	return root
}

type ListNode struct {
	Val  int
	Next *ListNode
}
