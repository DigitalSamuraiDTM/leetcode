package main

func main() {
	e1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	e2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	e3 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	e4 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	e5 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	out(removeNthFromEnd(e1, 1))
	out(removeNthFromEnd(e2, 2))
	out(removeNthFromEnd(e3, 3))
	out(removeNthFromEnd(e4, 4))
	out(removeNthFromEnd(e5, 5))

}

func out(head *ListNode) {
	c := head
	for true {
		if c == nil {
			println("FINISH")
			return
		}
		println(c.Val)
		c = c.Next
	}
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// forward for last node
	var removedNode = head
	var lastNode = head
	for i := 1; i < n; i++ {
		lastNode = lastNode.Next
	}
	// case: need remove first node
	if lastNode.Next == nil {
		return removedNode.Next
	}
	// find real removed node
	parentNode := removedNode
	for lastNode.Next != nil {
		lastNode = lastNode.Next
		parentNode = removedNode
		removedNode = removedNode.Next
	}
	// case: need remove LAST node
	if parentNode.Next.Next == nil {
		parentNode.Next = nil
	} else {
		parentNode.Next = parentNode.Next.Next
	}
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
