package main

func main() {
	e1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}

	out(swapPairs(e1))

	e2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	out(swapPairs(e2))

	e3 := &ListNode{
		Val:  1,
		Next: nil,
	}

	out(swapPairs(e3))
}

func out(head *ListNode) {
	ass := head
	for true {
		if ass == nil {
			println("FINISH")
			return
		}
		println(ass.Val)
		ass = ass.Next
	}
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	fNode := head
	sNode := head.Next
	for true {
		if sNode != nil && fNode != nil {
			// swap
			fNode.Val, sNode.Val = sNode.Val, fNode.Val
			fNode = sNode.Next
			if fNode == nil {
				break
			}
			sNode = fNode.Next
		} else {
			break
		}
	}
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
