package main

import (
	"slices"
)

func main() {
	examples := []string{"e", "2", "0089", "-0.1", "+3.14", "4.", "-.9", "2e10", "-90E3", "3e+7", "+6e-1", "53.5e93", "-123.456e789"}
	for _, example := range examples {
		out := isNumber(example)
		println(out)
	}
}

var symbols = []string{"+", "-"}
var numbers = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
var dot = "."
var exponenta = []string{"E", "e"}

// реализуем детерменированный конечный автомат для валидации строк
var validatorAutomat = map[int]func(char uint8) *int{
	0: func(char uint8) *int {
		if slices.Contains(symbols, string(char)) {
			x1 := 1
			return &x1
		}
		if string(char) == dot {
			x3 := 3
			return &x3
		}
		if slices.Contains(numbers, string(char)) {
			x2 := 2
			return &x2
		}
		return nil
	},
	1: func(char uint8) *int {
		if slices.Contains(numbers, string(char)) {
			x2 := 2
			return &x2
		}
		if string(char) == dot {
			x3 := 3
			return &x3
		}
		return nil
	},
	2: func(char uint8) *int {
		if slices.Contains(numbers, string(char)) {
			x2 := 2
			return &x2
		}
		if string(char) == dot {
			x8 := 8
			return &x8
		}
		if slices.Contains(exponenta, string(char)) {
			x5 := 5
			return &x5
		}
		return nil
	},
	3: func(char uint8) *int {
		if slices.Contains(numbers, string(char)) {
			x4 := 4
			return &x4
		}
		return nil
	},
	4: func(char uint8) *int {
		if slices.Contains(numbers, string(char)) {
			x4 := 4
			return &x4
		}
		if slices.Contains(exponenta, string(char)) {
			x5 := 5
			return &x5
		}
		return nil
	},
	5: func(char uint8) *int {
		if slices.Contains(numbers, string(char)) {
			x7 := 7
			return &x7
		}
		if slices.Contains(symbols, string(char)) {
			x6 := 6
			return &x6
		}
		return nil
	},
	6: func(char uint8) *int {
		if slices.Contains(numbers, string(char)) {
			x7 := 7
			return &x7
		}
		return nil
	},
	7: func(char uint8) *int {
		if slices.Contains(numbers, string(char)) {
			x7 := 7
			return &x7
		}
		return nil
	},
	8: func(char uint8) *int {
		if slices.Contains(numbers, string(char)) {
			x4 := 4
			return &x4
		}
		if slices.Contains(exponenta, string(char)) {
			x5 := 5
			return &x5
		}
		return nil
	},
}

func isNumber(s string) bool {
	automatState := 0
	for i := range s {
		char := s[i]
		nextAutomatState := validatorAutomat[automatState](char)
		if nextAutomatState == nil {
			return false
		} else {
			automatState = *nextAutomatState
		}
	}
	return automatState == 2 || automatState == 4 || automatState == 7 || automatState == 8
}
