package main

func main() {
	e1 := []int{2, 2, 1}
	println(singleNumber(e1))

	e2 := []int{4, 1, 2, 1, 2}
	println(singleNumber(e2))

	e3 := []int{1}
	println(singleNumber(e3))

	e4 := []int{1, 2, 3, 4, 3, 2, 1}
	println(singleNumber(e4))
}

func singleNumber(nums []int) int {
	checkedMap := make(map[int]struct{})
	for _, value := range nums {
		_, isExist := checkedMap[value]
		if isExist {
			delete(checkedMap, value)
		} else {
			checkedMap[value] = struct{}{}
		}
	}
	out := 0
	for key, _ := range checkedMap {
		out = key
	}
	return out
}
