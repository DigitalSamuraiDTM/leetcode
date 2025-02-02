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
	out(reverseList(e1))

	e2 := &ListNode{
		Val:  1,
		Next: nil,
	}
	out(reverseList(e2))

	out(reverseList(nil))

}

func out(listNode *ListNode) {
	c := listNode
	for true {
		if c == nil {
			println("FINISH")
			return
		} else {
			println(c.Val)
			c = c.Next
		}
	}
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	node := &ListNode{
		Val: head.Val,
	}
	nextNode := head.Next

	for true {
		if nextNode == nil {
			return node
		} else {
			node = &ListNode{
				Val:  nextNode.Val,
				Next: node,
			}
			nextNode = nextNode.Next
		}
	}
	return nil
}

type ListNode struct {
	Val  int
	Next *ListNode
}
