package main

func main() {
	println(missingNumber2([]int{3, 0, 1}))
	println(missingNumber2([]int{0, 1}))
	println(missingNumber2([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
	println(missingNumber2([]int{0, 1, 2, 3, 4, 6, 7, 8, 9, 10, 11, 12}))
	println(missingNumber2([]int{1, 2, 3, 4}))
}

func missingNumber(nums []int) int {
	validatorMap := make(map[int]bool, len(nums))
	for _, num := range nums {
		validatorMap[num] = true
	}

	for index := range nums {
		if validatorMap[index] == false {
			return index
		}
	}
	return len(nums)
}

func missingNumber2(nums []int) int {
	out := len(nums)
	for index := range nums {
		out = out + index - nums[index]
	}
	return out
}
