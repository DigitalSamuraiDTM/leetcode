package main

func main() {
	e1 := &ListNode{
		Val:  1,
		Next: nil,
	}
	e1.Next = e1
	println(hasCycle(e1))
	println(hasCycleConstTime(e1))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	nodeMarkerMap := map[*ListNode]bool{}
	node := head
	for true {
		nodeMarkerMap[node] = true
		if node.Next == nil {
			break
		} else {
			if nodeMarkerMap[node.Next] == true {
				return true
			} else {
				node = node.Next
			}
		}
	}
	return false
}

// CONST MEMORY
func hasCycleConstTime(head *ListNode) bool {
	if head == nil {
		return false
	}
	// in task we has node limits and it is condition for cycle break
	taskLimit := 10000
	index := 0
	node := head
	for index <= taskLimit {
		if node.Next == nil {
			return false
		}
		node = node.Next
		index++
	}
	return true
}
