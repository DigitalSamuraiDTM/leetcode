package main

import "math"

func main() {
	e1x1 := [][]int{
		{1},
	}
	e2x2 := [][]int{
		{1, 2},
		{3, 4},
	}
	e3x3 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	e4x4 := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	e5x5 := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25},
	}
	e6x6 := [][]int{
		{1, 2, 3, 4, 5, 6},
		{7, 8, 9, 10, 12, 13},
		{14, 15, 16, 17, 18, 19},
		{20, 21, 22, 23, 24, 25},
		{26, 27, 28, 29, 30, 31},
		{32, 33, 34, 35, 36, 37},
	}
	rotate(e1x1)
	rotate(e2x2)
	rotate(e3x3)
	rotate(e4x4)
	rotate(e5x5)
	rotate(e6x6)

	out(e1x1)
	println("NEXT")
	out(e2x2)
	println("NEXT")
	out(e3x3)
	println("NEXT")
	out(e4x4)
	println("NEXT")
	out(e5x5)
	println("NEXT")
	out(e6x6)
	println("NEXT")
}
func out(matrix [][]int) {
	i := 0
	for i < len(matrix) {
		j := 0
		for j < len(matrix[i]) {
			print(matrix[i][j])
			print(" ")
			j++
		}
		println("")
		i++
	}
}

func rotate(matrix [][]int) {
	r := 0
	n := len(matrix[0])
	for r != int(math.Ceil(float64(n-1)/2)) {
		iter := n - 1 - r*2
		i := 0
		for i < iter {
			// get
			el1 := matrix[r][r+i]
			el2 := matrix[r+i][n-1-r]
			el3 := matrix[n-1-r][n-1-r-i]
			el4 := matrix[n-1-r-i][r]
			// rotate!!
			matrix[r+i][n-1-r] = el1
			matrix[n-1-r][n-1-r-i] = el2
			matrix[n-1-r-i][r] = el3
			matrix[r][r+i] = el4
			i++
		}
		r++
	}
}
