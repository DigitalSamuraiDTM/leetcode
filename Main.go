package main

func main() {
	out(letterCombinations("23"))
	out(letterCombinations(""))
	out(letterCombinations("7"))
	out(letterCombinations("237"))
	out(letterCombinations("149"))
}

func out(out []string) {
	for _, value := range out {
		print(value + " ")
	}
	println("OUT")
}

var numbers = map[uint8][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return make([]string, 0)
	}
	if len(digits) == 1 {
		return numbers[digits[0]]
	}
	outSize := 1
	for index, _ := range digits {
		outSize *= len(numbers[digits[index]])
	}
	outArray := make([]string, outSize)
	currentPosition := 0
	reCursed(outArray, &currentPosition, digits, "")
	return outArray
}

func reCursed(outArray []string, currentPosition *int, digits string, currentSubstring string) {
	if len(digits) == 1 {
		for _, value := range numbers[digits[0]] {
			outArray[*currentPosition] = currentSubstring + value
			*currentPosition++
		}
	} else {
		for _, value := range numbers[digits[0]] {
			reCursed(outArray, currentPosition, digits[1:], currentSubstring+value)
		}
	}
}
