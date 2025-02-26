package main

func main() {
	println(addBinary("1", "111"))
	println(addBinary("1010", "1011"))
	println(addBinary("11", "1"))
	println(addBinary("111", "111"))
}

func addBinary(a string, b string) string {
	aArr := []rune(a)
	bArr := []rune(b)
	index := 0
	out := ""
	hasTail := false
	for index < len(aArr) && index < len(bArr) {
		aSymbol := aArr[len(aArr)-1-index]
		bSymbol := bArr[len(bArr)-1-index]
		if aSymbol == '1' && bSymbol == '1' {
			if hasTail {
				out = "1" + out
				hasTail = true
			} else {
				hasTail = true
				out = "0" + out
			}
		} else if aSymbol == '0' && bSymbol == '0' {
			if hasTail {
				hasTail = false
				out = "1" + out
			} else {
				out = "0" + out
			}
		} else {
			if hasTail {
				out = "0" + out
				hasTail = true
			} else {
				out = "1" + out
			}
		}
		index++
	}
	if len(aArr) > index {
		for index < len(aArr) {
			aSymbol := aArr[len(aArr)-1-index]
			if aSymbol == '0' && hasTail {
				out = "1" + out
				hasTail = false
			} else if aSymbol == '1' && hasTail {
				out = "0" + out
			} else {
				out = string(aSymbol) + out
			}
			index++
		}
	} else if len(bArr) > index {
		for index < len(bArr) {
			bSymbol := bArr[len(bArr)-1-index]
			if bSymbol == '0' && hasTail {
				out = "1" + out
				hasTail = false
			} else if bSymbol == '1' && hasTail {
				out = "0" + out
			} else {
				out = string(bSymbol) + out
			}
			index++
		}
	}

	if hasTail {
		out = "1" + out
	}
	return out
}
