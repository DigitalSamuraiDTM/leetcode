package main

import (
	"slices"
)

func main() {
	println(myAtoi("42"))
	println(myAtoi("-042"))
	println(myAtoi("1337c0d3"))
	println(myAtoi("0-1"))
	println(myAtoi("-000001"))
	println(myAtoi("words and 987"))
	println(myAtoi("-999999999999999999999999999999999999999999"))
	println(myAtoi("2147483649"))           // 2147483647
	println(myAtoi("2147484647"))           // 2147483647
	println(myAtoi("-2147483649"))          // -2147483648
	println(myAtoi("-0000000002147483648")) // -2147483648
	println(myAtoi("2147483647"))           // 2147483647
	println(myAtoi("21474836475"))          // 2147483647
	println(myAtoi("21474836460"))          // 2147483647
	println(myAtoi("1095502006p8"))         // 1095502006
}

type ValidatorAutomat struct {
	index           int
	validatedNumber int
	isNegative      bool
	canOverlimit    bool
	isOverlimit     bool
	state           int
	states          map[int]func(char uint8, automat *ValidatorAutomat) bool
}

func (automat *ValidatorAutomat) Next(char uint8) bool {

	var isOk = automat.states[automat.state](char, automat)
	return isOk
}

func (automat *ValidatorAutomat) AddNumber(number int) {
	if automat.index == 9 {
		automat.isOverlimit = automat.validatedNumber > 214748364 ||
			automat.validatedNumber == 214748364 && automat.isNegative && number >= 8 ||
			automat.validatedNumber == 214748364 && !automat.isNegative && number >= 7
	}
	if automat.index >= 10 {
		automat.isOverlimit = true
	}
	if automat.isOverlimit {
		automat.index++
		return
	}
	automat.index++
	automat.validatedNumber = automat.validatedNumber*10 + number
}

func (automat *ValidatorAutomat) MakeNegative() {
	automat.isNegative = true
}

func (automat *ValidatorAutomat) GetValidatedNumber() int {
	if automat.isOverlimit {
		if automat.isNegative {
			return -2147483648
		} else {
			return 2147483647
		}
	}
	var neg = 1
	if automat.isNegative {
		neg = -1
	}
	return automat.validatedNumber * neg
}

func InitAutomat() ValidatorAutomat {
	var numbers = []uint8{'1', '2', '3', '4', '5', '6', '7', '8', '9'}
	automatStates := map[int]func(char uint8, automat *ValidatorAutomat) bool{
		0: func(char uint8, automat *ValidatorAutomat) bool {
			if char == ' ' {
				return true
			}
			if char == '-' {
				automat.state = 1
				automat.MakeNegative()
				return true
			}
			if char == '0' || char == '+' {
				automat.state = 1
				return true
			}
			if slices.Contains(numbers, char) {
				automat.state = 3
				automat.AddNumber(int(char) - 48)
				return true
			}
			return false
		},
		1: func(char uint8, automat *ValidatorAutomat) bool {
			if char == '0' {
				return true
			}
			if slices.Contains(numbers, char) {
				automat.state = 3
				automat.AddNumber(int(char) - 48)
				return true
			}
			return false
		},
		3: func(char uint8, automat *ValidatorAutomat) bool {
			if slices.Contains(numbers, char) || char == '0' {
				automat.state = 4
				automat.AddNumber(int(char) - 48)
				return true
			}
			return false
		},
		4: func(char uint8, automat *ValidatorAutomat) bool {
			if slices.Contains(numbers, char) || char == '0' {
				automat.AddNumber(int(char) - 48)
				return true
			}
			return false
		},
	}
	return ValidatorAutomat{
		states:          automatStates,
		validatedNumber: 0,
		state:           0,
	}
}

func myAtoi(s string) int {

	automat := InitAutomat()
	for _, char := range s {
		isOk := automat.Next(uint8(char))
		if !isOk {
			break
		}
	}
	return automat.GetValidatedNumber()
}
