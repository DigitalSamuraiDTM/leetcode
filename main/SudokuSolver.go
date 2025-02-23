package main

import (
	"fmt"
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
	Num    int
	Row    int
	Column int
}

func (s *SudokuSolver) SolveIt() {
	s.inputValidation()
	s.setupAllowedNumber()
	s.solveRecursively()
}

func (s *SudokuSolver) solveRecursively() bool {

	// find available places (if place was not found -> sudoku invalid and we need reverse)
	number, places := s.findRandomAvailablePlace()
	println(fmt.Sprintf("PLACES: %d, EMPTY: %d", len(places), s.emptyPlaces))
	if number == nil {
		return false
	}
	for _, place := range places {
		// try to setup
		s.setNumber(*number, place.Row, place.Column)

		// validate
		setupedPlaces := s.setupAllowedNumber()

		if s.isSudokuSolved() {
			return true
		}
		if s.solveRecursively() {
			return true
		} else {
			s.unsetNumber(*number, place.Row, place.Column)
			for _, setupedPlace := range setupedPlaces {
				s.unsetNumber(setupedPlace.Num, setupedPlace.Row, setupedPlace.Column)
			}
		}
	}

	// repeat
	return s.isSudokuSolved()
}

func (s *SudokuSolver) findAllAvailablePlaces() {

}

// return num and valid places for it. if num is null -> sudoku invalid
func (s *SudokuSolver) findRandomAvailablePlace() (num *int, places []Place) {
	for ceilNumber, value := range s.validCeilMatrix {
		for i := 1; i < 10; i++ {
			if value[i] == true {
				// num exist at this ceil
				continue
			} else {
				// get available places
				places = s.getAvailablePlacesInCeil(ceilNumber, i)
				// if available places was founded then return
				if len(places) != 0 {
					num = &i
					return
				}
			}
		}
	}
	// invalid sudoku
	return nil, nil
}

func (s *SudokuSolver) setupAllowedNumber() (setupedList []Place) {
	number := 1
	for true {
		isNeedRestart := false
		for ceilNumber, value := range s.validCeilMatrix {
			if value[number] == false {
				availablePlaces := s.getAvailablePlacesInCeil(ceilNumber, number)
				// if only one condition in ceil
				if len(availablePlaces) == 1 {
					s.setNumber(number, availablePlaces[0].Row, availablePlaces[0].Column)
					setupedList = append(setupedList, availablePlaces...)
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
					availablePlaces = append(availablePlaces, Place{Num: number, Row: row, Column: col})
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
