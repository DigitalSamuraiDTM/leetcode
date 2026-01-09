package main

func main() {
	e1 := []int{0, 1, 0, 3, 12}
	moveZeroes(e1)
	printResult(e1)

	e2 := []int{0, 1, 0, 3, 0, 12, 0, 0, 16, 0, 17, 18, 19, 0, 0, 20, 0, 21, 0}
	moveZeroes(e2)
	printResult(e2)
}

func printResult(nums []int) {
	for _, num := range nums {
		println(num)
	}
	println(" ")
}

func moveZeroes(nums []int) {
	availablePlace := 0
	for i, num := range nums {
		if num != 0 {
			nums[availablePlace], nums[i] = nums[i], nums[availablePlace]
			availablePlace++
		}
	}

}
