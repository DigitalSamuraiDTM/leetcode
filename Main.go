package main

func main() {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	list2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	case1 := mergeTwoLists(list1, list2)
	out(case1)
}

func out(listNode *ListNode) {
	if listNode == nil {
		return
	}
	println(listNode.Val)
	out(listNode.Next)
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// finish condition
	if list1 == nil && list2 == nil {
		return nil
	}
	// fast forward first list
	if list1 != nil && list2 == nil {
		return &ListNode{
			Val:  list1.Val,
			Next: mergeTwoLists(list1.Next, nil),
		}
	}
	// fast forward second list
	if list1 == nil && list2 != nil {
		return &ListNode{
			Val:  list2.Val,
			Next: mergeTwoLists(list2.Next, nil),
		}
	}
	// sorting
	if list1.Val < list2.Val {
		return &ListNode{
			Val:  list1.Val,
			Next: mergeTwoLists(list1.Next, list2),
		}
	} else {
		return &ListNode{
			Val:  list2.Val,
			Next: mergeTwoLists(list1, list2.Next),
		}
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}
