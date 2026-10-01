package main

func main() {
	println(isAnagram("anagram", "nagaram"))
	println(isAnagram("rat", "car"))
	println(isAnagram("", ""))
}

func isAnagram(s string, t string) bool {

	anagramMap := map[uint8]int{}
	for i := range s {
		curr := s[i]
		anagramMap[curr] += 1
	}
	// TODO в 1 цикл. Один прибавляет, второй отнимает
	for i := range t {
		curr := t[i]
		anagramMap[curr] -= 1
		if anagramMap[curr] < 0 {
			return false
		}
	}
	for u := range anagramMap {
		if anagramMap[u] != 0 {
			return false
		}
	}
	return true
}
