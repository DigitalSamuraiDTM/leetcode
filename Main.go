package main

func main() {
	out(generate(30))
}

func out(arr [][]int) {
	for i, _ := range arr {
		for _, value := range arr[i] {
			print(value)
			print(", ")
		}
		println()
	}
}

func generate(numRows int) [][]int {
	triangle := make([][]int, numRows)
	if numRows >= 1 {
		triangle[0] = []int{1}
	}
	if numRows >= 2 {
		triangle[1] = []int{1, 1}
	}
	indexator := 2
	for indexator < numRows {
		topFacade := triangle[indexator-1]
		lowerFacade := []int{1}
		for i := 0; i < len(topFacade)-1; i++ {
			lowerFacade = append(lowerFacade, topFacade[i]+topFacade[i+1])
		}
		lowerFacade = append(lowerFacade, 1)
		triangle[indexator] = lowerFacade
		indexator++
	}
	return triangle
}
