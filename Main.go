package main

func main() {
	e1 := []int{1, 2, 3, 4, 5}
	e2 := []int{9, 9, 9}
	e3 := []int{9, 8, 9}
	e4 := []int{1}
	e5 := []int{1, 9, 9, 9, 9, 9, 9, 9, 9, 9}
	e6 := []int{9, 0, 0}
	out(plusOne(e1))
	out(plusOne(e2))
	out(plusOne(e3))
	out(plusOne(e4))
	out(plusOne(e5))
	out(plusOne(e6))
}

func out(out []int) {
	for _, value := range out {
		print(value)
		print(" ")
	}
	println()
	print()
}

func plusOne(digits []int) []int {
	if digits[len(digits)-1] != 9 {
		digits[len(digits)-1]++
		return digits
	}
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i]++
			return digits
		}
	}

	digits[0] = 0
	out := append([]int{1}, digits...)
	return out
}
