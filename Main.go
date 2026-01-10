package main

func main() {
	//e1 := []int{1, 2, 3, 0, 0, 0}
	//merge(e1, 3, []int{2, 5, 6}, 3)
	//printResult(e1)
	//
	//e2 := []int{1}
	//merge(e2, 1, []int{}, 0)
	//printResult(e2)
	//
	//e3 := []int{0}
	//merge(e3, 0, []int{1}, 1)
	//printResult(e3)
	//
	//e4 := []int{1, 2, 3, 6, 8, 9, 0, 0, 0, 0, 0}
	//merge(e4, 6, []int{2, 2, 6, 6, 10}, 5)
	//printResult(e4) // 1 2 2 2 3 6 6 6 8 9 10
	//
	//e5 := []int{2, 0}
	//merge(e5, 1, []int{1}, 1)
	//printResult(e5)
	//
	//e6 := []int{1, 2, 4, 5, 6, 0}
	//merge(e6, 5, []int{3}, 1)
	//printResult(e6)

	//e7 := []int{-1, 0, 0, 3, 3, 3, 0, 0, 0}
	//merge(e7, 6, []int{1, 2, 2}, 3)
	//printResult(e7)

	e8 := []int{0, 0, 3, 0, 0, 0, 0, 0, 0}
	merge(e8, 3, []int{-1, 1, 1, 1, 2, 3}, 6)
	printResult(e8)
}

func printResult(result []int) {
	for _, value := range result {
		println(value)
	}
	println(" ")
}

// O(m+n)
func merge(nums1 []int, m int, nums2 []int, n int) {
	nums2Index := 0
	bufferIndex := 0
	if n == 0 {
		return
	}
	buffer := make([]int, m)
	bufferSaturatorIndex := 0
	for i := range nums1 {
		if i+1 > m {
			if nums2Index == len(nums2) {
				nums1[i] = buffer[bufferIndex]
				bufferIndex++
			} else if bufferIndex < bufferSaturatorIndex && buffer[bufferIndex] < nums2[nums2Index] {
				nums1[i] = buffer[bufferIndex]
				bufferIndex++
			} else {
				nums1[i] = nums2[nums2Index]
				nums2Index++
			}
		} else {
			if nums2Index == len(nums2) || bufferIndex < bufferSaturatorIndex && buffer[bufferIndex] < nums1[i] && buffer[bufferIndex] < nums2[nums2Index] {
				buffer[bufferSaturatorIndex] = nums1[i]
				bufferSaturatorIndex++
				nums1[i] = buffer[bufferIndex]
				bufferIndex++
			} else if nums2[nums2Index] < nums1[i] {
				buffer[bufferSaturatorIndex] = nums1[i]
				bufferSaturatorIndex++
				nums1[i] = nums2[nums2Index]
				nums2Index++
			}
		}
	}
}
