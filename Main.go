package main

func main() {
	var e1 = []int{3, 2, 3}
	println(majorityElement(e1))
	var e2 = []int{2, 2, 1, 1, 1, 2, 2}
	println(majorityElement(e2))
	var e3 = []int{2, 2, 2, 1, 2, 1, 1}
	println(majorityElement(e3))
}

//func majorityElement(nums []int) int {
//	var mapElements = make(map[int]int)
//	var maxElement = 0
//	for i := range nums {
//		element := nums[i]
//		mapElements[element]++
//		if mapElements[element] > maxElement {
//			maxElement = element
//		}
//	}
//	return maxElement
//}

func majorityElement(nums []int) int {
	var currentNum = 0
	var counter = 0

	for i := range nums {
		if counter == 0 {
			currentNum = nums[i]
		}
		if currentNum == nums[i] {
			counter++
		} else {
			counter--
		}
	}
	return currentNum
}
