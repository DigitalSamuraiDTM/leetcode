package main

func main() {
	e1 := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	println(isValidSudoku(e1))
}

func isValidSudoku(board [][]byte) bool {

	validRowMatrix := make([]map[byte]bool, 9)
	validColMatrdix := make([]map[byte]bool, 9)
	validCeilMatrix := make([]map[byte]bool, 9)

	for rowIndex, row := range board {
		for columnIndex, item := range row {
			if item != '.' {
				// by default map is nil
				if validRowMatrix[rowIndex] == nil {
					validRowMatrix[rowIndex] = map[byte]bool{}
				}
				// validate row
				if validRowMatrix[rowIndex][item-1] == true {
					return false
				} else {
					validRowMatrix[rowIndex][item-1] = true
				}
				// by default map is nil
				if validColMatrdix[columnIndex] == nil {
					validColMatrdix[columnIndex] = map[byte]bool{}
				}
				// validate column
				if validColMatrdix[columnIndex][item-1] == true {
					return false
				} else {
					validColMatrdix[columnIndex][item-1] = true
				}
				ceilIndex := rowIndex/3*3 + columnIndex/3
				if validCeilMatrix[ceilIndex] == nil {
					validCeilMatrix[ceilIndex] = map[byte]bool{}
				}
				// validate ceil
				if validCeilMatrix[ceilIndex][item-1] == true {
					return false
				} else {
					validCeilMatrix[ceilIndex][item-1] = true
				}
			}
		}
	}
	return true
}
