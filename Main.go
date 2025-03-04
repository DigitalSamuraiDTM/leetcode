package main

func main() {
	e1 := &ListNode{
		Val: 6,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 5,
							Next: &ListNode{
								Val:  6,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}
	out(removeElements(e1, 6))
}

func out(node *ListNode) {
	v := node
	for true {
		if v == nil {
			println("FINISH")
			return
		}
		println(v.Val)
		v = v.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	v := head
	var newRoot *ListNode
	var pointer *ListNode
	for v != nil {
		if v.Val != val {
			if newRoot == nil {
				newRoot = &ListNode{
					Val:  v.Val,
					Next: nil,
				}
				pointer = newRoot
			} else {
				pointer.Next = &ListNode{
					Val:  v.Val,
					Next: nil,
				}
				pointer = pointer.Next
			}
		}
		v = v.Next
	}
	return newRoot
}
