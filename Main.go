package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	//var intersecion1 = &ListNode{
	//	Val: 8,
	//	Next: &ListNode{
	//		Val: 4,
	//		Next: &ListNode{
	//			Val:  5,
	//			Next: nil,
	//		},
	//	},
	//}
	//var a1 = &ListNode{
	//	Val: 4,
	//	Next: &ListNode{
	//		Val:  1,
	//		Next: intersecion1,
	//	},
	//}
	//var b1 = &ListNode{
	//	Val: 5,
	//	Next: &ListNode{
	//		Val: 6,
	//		Next: &ListNode{
	//			Val:  1,
	//			Next: intersecion1,
	//		},
	//	},
	//}
	//var out1 = getIntersectionNode(a1, b1)
	//output(out1)

	var intersection2 = &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  4,
			Next: nil,
		},
	}
	var a2 = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val:  1,
				Next: intersection2,
			},
		},
	}
	var b2 = &ListNode{
		Val:  3,
		Next: intersection2,
	}
	var out2 = getIntersectionNode(a2, b2)
	output(out2)
}

func output(node *ListNode) {
	for node != nil {
		println(node.Val)
		node = node.Next
	}
	println("END")
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var a = headA
	var b = headB
	var markerMap = make(map[*ListNode]bool)
	for {
		if a == nil && b == nil {
			break
		}
		if a != nil {
			if markerMap[a] == true {
				return a
			} else {
				markerMap[a] = true
			}
			a = a.Next
		}
		if b != nil {
			if markerMap[b] == true {
				return b
			} else {
				markerMap[b] = true
			}
			b = b.Next
		}
	}
	return nil
}
