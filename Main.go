package main

import "strings"

func main() {
	//printResult(fullJustify([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16))
	//println(" ")
	//printResult(fullJustify([]string{"What", "must", "be", "acknowledgment", "shall", "be"}, 16))
	//println(" ")
	//printResult(fullJustify([]string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}, 20))
	//println(" ")
	printResult(fullJustify([]string{"The", "important", "thing", "is", "not", "to", "stop", "questioning.", "Curiosity", "has", "its", "own", "reason", "for", "existing."}, 17))

}

func printResult(list []string) {
	for _, s := range list {
		println(s)
	}
}

func fullJustify(words []string, maxWidth int) []string {
	limitWidthCounter := 0
	wordCounter := 0
	justificatedIndex := 0
	outResult := make([]string, 0)
	for i := range words {
		word := words[i]
		if limitWidthCounter+len(word) > maxWidth {
			// begin middle justification
			if wordCounter == 1 {
				outResult = append(outResult, leftJustifyString([]string{words[i-1]}, maxWidth))
				justificatedIndex = i
				wordCounter = 1
				limitWidthCounter = len(word) + 1
				continue
			}
			spacePlaceCount := wordCounter - 1              // 3
			symbolsCount := limitWidthCounter - wordCounter // 19 - 4 = 15
			spacesCount := maxWidth - symbolsCount          // 5
			// сколько надо поставить обязательных пробелов
			requiredSpaces := spacesCount / spacePlaceCount // 1
			// сколько экстра пробелов
			extraSpaces := spacesCount % spacePlaceCount // 2
			slice := words[justificatedIndex:i]
			if extraSpaces == 0 {
				outString := strings.Join(slice, strings.Repeat(" ", requiredSpaces))
				outResult = append(outResult, outString)
			} else {
				outString := ""
				for j := range slice {
					extraSpace := ""
					if extraSpaces > 0 {
						extraSpace = " "
						extraSpaces--
					}
					if len(slice) == j+1 {
						outString += slice[j]
						continue
					}
					outString += slice[j] + strings.Repeat(" ", requiredSpaces) + extraSpace
				}
				outResult = append(outResult, outString)
			}

			justificatedIndex = i
			wordCounter = 1
			limitWidthCounter = len(word) + 1
		} else {
			wordCounter++
			limitWidthCounter += len(word) + 1
		}
	}
	// at last left
	if len(words) > justificatedIndex {
		slice := words[justificatedIndex:]
		outResult = append(outResult, leftJustifyString(slice, maxWidth))
	}
	return outResult
}

func leftJustifyString(slice []string, width int) string {
	joined := strings.Join(slice, " ")
	joined += strings.Repeat(" ", width-len(joined))
	return joined
}
