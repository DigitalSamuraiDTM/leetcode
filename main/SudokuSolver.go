package main

import (
	"strconv"
)

type SudokuSolver struct {
	board           *[][]byte
	validCeilMatrix []map[int]bool
	validColMatrix  []map[int]bool
	validRowMatrix  []map[int]bool

	emptyPlaces int
}

type Place struct {
	Row    int
	Column int
}

func (s *SudokuSolver) SolveIt() {
	s.inputValidation()
	availablePlaces := s.setupAllowedNumber()
	if s.isSudokuSolved() {
		return
	}
	s.solveRecursively(availablePlaces, -1, 1)
}

func (s *SudokuSolver) solveRecursively(availablePlaces map[int]map[int][]Place, ceil int, number int) bool {

	newCeil := ceil
	newNumber := number
	places := availablePlaces[newCeil][newNumber]
	for true {
		if newCeil == 8 {
			newCeil = 0
			newNumber++
		} else {
			newCeil++
		}
		places = availablePlaces[newCeil][newNumber]
		if places != nil && len(places) != 0 {
			break
		}
	}

	for _, value := range places {
		if (*s.board)[value.Row][value.Column] != '.' || !s.isValidPlace(value.Row, value.Column, newNumber) {
			continue
		}
		s.setNumber(newNumber, value.Row, value.Column)
		if s.isSudokuSolved() {
			return true
		}
		isSolved := s.solveRecursively(availablePlaces, newCeil, newNumber)
		if isSolved {
			return true
		} else {
			//println(fmt.Sprintf("UNSET IN ROW: %d, COL: %d, NUM: %d", value.Row, value.Column, newNumber))
			s.unsetNumber(newNumber, value.Row, value.Column)
		}
	}

	return s.isSudokuSolved()
}

func (s *SudokuSolver) setupAllowedNumber() (setupedList map[int]map[int][]Place) {
	setupedList = make(map[int]map[int][]Place)
	for i := 0; i < 9; i++ {
		setupedList[i] = make(map[int][]Place)
	}
	number := 1
	for true {
		isNeedRestart := false
		for ceilNumber, value := range s.validCeilMatrix {
			if value[number] == false {
				availablePlaces := s.getAvailablePlacesInCeil(ceilNumber, number)
				// if only one condition in ceil
				if len(availablePlaces) == 1 {
					s.setNumber(number, availablePlaces[0].Row, availablePlaces[0].Column)
					isNeedRestart = true
					break
				}
			}
		}
		if !isNeedRestart {
			number++
		} else {
			isNeedRestart = false
			number = 1
		}
		if number == 10 {
			// finish setup allowed places
			break
		}
	}
	for n := 1; n < 10; n++ {
		for ceilNumber, value := range s.validCeilMatrix {
			if value[n] == false {
				availablePlaces := s.getAvailablePlacesInCeil(ceilNumber, n)
				// if only one condition in ceil
				if len(availablePlaces) > 1 {
					setupedList[ceilNumber][n] = availablePlaces
				}
			}
		}
	}
	return
}

func (s *SudokuSolver) getAvailablePlacesInCeil(numberOfCeil int,
	number int,
) (availablePlaces []Place) {
	// try setup if we can
	ceilRow := (numberOfCeil / 3) * 3
	ceilCol := (numberOfCeil % 3) * 3
	for row := ceilRow; row < ceilRow+3; row++ {
		for col := ceilCol; col < ceilCol+3; col++ {
			if (*s.board)[row][col] == '.' {
				// assumption, can we setup num at this place?
				isValid := s.isValidPlace(row, col, number)
				if isValid {
					availablePlaces = append(availablePlaces, Place{Row: row, Column: col})
				}
			}
		}
	}
	return
}

func (s *SudokuSolver) setNumber(number int, row int, col int) {
	//println(fmt.Sprintf("N: %d, R: %d, C: %d", number, row, col))
	s.validCeilMatrix[row/3*3+col/3][number] = true
	s.validRowMatrix[row][number] = true
	s.validColMatrix[col][number] = true
	(*s.board)[row][col] = ([]byte(strconv.Itoa(number)))[0]
	s.emptyPlaces--
}

func (s *SudokuSolver) unsetNumber(number int, row int, col int) {
	s.validCeilMatrix[row/3*3+col/3][number] = false
	s.validRowMatrix[row][number] = false
	s.validColMatrix[col][number] = false
	(*s.board)[row][col] = '.'
	s.emptyPlaces++
}

func (s *SudokuSolver) isSudokuSolved() bool {
	return s.emptyPlaces == 0
}

func (s *SudokuSolver) isValidPlace(
	row int,
	col int,
	number int,
) bool {
	return s.validRowMatrix[row][number] == false &&
		s.validColMatrix[col][number] == false &&
		s.validCeilMatrix[row/3*3+col/3][number] == false
}

func (s *SudokuSolver) inputValidation() {
	for i := 0; i < 9; i++ {
		s.validCeilMatrix[i] = map[int]bool{}
		s.validRowMatrix[i] = map[int]bool{}
		s.validColMatrix[i] = map[int]bool{}
	}

	for rowIndex, row := range *s.board {
		for columnIndex, item := range row {
			if item != '.' {
				// by default map is nil
				if s.validRowMatrix[rowIndex] == nil {
					s.validRowMatrix[rowIndex] = map[int]bool{}
				}
				// validate row
				if s.validRowMatrix[rowIndex][int(item)] == true {
					continue
				} else {
					n, _ := strconv.Atoi(string(item))
					s.validRowMatrix[rowIndex][n] = true
				}
				// by default map is nil
				if s.validColMatrix[columnIndex] == nil {
					s.validColMatrix[columnIndex] = map[int]bool{}
				}
				// validate column
				if s.validColMatrix[columnIndex][int(item)] == true {
					continue
				} else {
					n, _ := strconv.Atoi(string(item))
					s.validColMatrix[columnIndex][n] = true
				}
				ceilIndex := rowIndex/3*3 + columnIndex/3
				if s.validCeilMatrix[ceilIndex] == nil {
					s.validCeilMatrix[ceilIndex] = map[int]bool{}
				}
				// validate ceil
				if s.validCeilMatrix[ceilIndex][int(item)] == true {
					continue
				} else {
					n, _ := strconv.Atoi(string(item))
					s.validCeilMatrix[ceilIndex][n] = true
				}
			} else {
				s.emptyPlaces++
			}
		}
	}
}
