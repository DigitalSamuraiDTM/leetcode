package main

func main() {
	e5 := []int{1, 2}
	val5 := 2
	out(e5, removeElement(e5, val5))

	e1 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val1 := 2
	out(e1, removeElement(e1, val1))

	e2 := []int{3, 2, 2, 3}
	val2 := 3
	out(e2, removeElement(e2, val2))

	e3 := []int{0, 0, 0, 0, 0}
	val3 := 0
	out(e3, removeElement(e3, val3))

	e4 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	val4 := 9
	out(e4, removeElement(e4, val4))

}

func out(nums []int, outVal int) {
	for _, value := range nums {
		print(value)
		print(" ")
	}
	println()
	print("OUT VAL: ")
	print(outVal)
	println()
}

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	tailIndex := len(nums) - 1
	headIndex := 0
	for headIndex != tailIndex {
		if nums[headIndex] == val {
			// swap!
			nums[headIndex], nums[tailIndex] = nums[tailIndex], nums[headIndex]
			tailIndex--
		} else {
			headIndex++
		}
	}
	if nums[headIndex] != val {
		headIndex++
	}
	return headIndex
}
