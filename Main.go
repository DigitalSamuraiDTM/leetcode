package main

import "unicode"

func main() {
	println(isPalindrome("0P"))
	println(isPalindrome("aaa"))
	println(isPalindrome(".,"))
	println(isPalindrome(".,."))
	println(isPalindrome("A man, a plan, a canal: Panama"))
	println(isPalindrome("race a car"))
	println(isPalindrome(" "))
}

func isPalindrome(s string) bool {
	leftIndex := 0
	rightIndex := len(s) - 1
	for rightIndex > leftIndex {
		l := rune(s[leftIndex])
		r := rune(s[rightIndex])
		isLL := unicode.IsLetter(l) || unicode.IsNumber(l)
		isRL := unicode.IsLetter(r) || unicode.IsNumber(r)
		if isLL && isRL {
			if unicode.ToLower(l) != unicode.ToLower(r) {
				return false
			}
			leftIndex++
			rightIndex--
		} else {
			if !isLL {
				leftIndex++
			}
			if !isRL {
				rightIndex--
			}
		}
	}

	return true
}
