package main

func main() {
	println(containsDuplicate([]int{1, 2, 3, 4}))                         // false
	println(containsDuplicate([]int{1, 2, 3, 4, 1}))                      // true
	println(containsDuplicate([]int{1, 2, 3, 4, -1, 5}))                  // false
	println(containsDuplicate([]int{1, 2, 3, 4, -1, -2, -3, -4, -5, -1})) // true
	println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))       // true
}

func containsDuplicate(nums []int) bool {
	checker := make(map[int]bool)
	for _, value := range nums {
		if checker[value] == true {
			return true
		} else {
			checker[value] = true
		}
	}
	return false
}
