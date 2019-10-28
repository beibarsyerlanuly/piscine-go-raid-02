package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
//	args := []string{".96.4...1", "1...6...4", "5.481.39.", "..795..43", ".3..8....", "4.5.23.18", ".1.63..59", ".59.7.83.", "..359...7"}
//	args := []string{"5........", "..31.9...", "96.8.7..4",".729..5.8","...4357.1",".1.....6.","7582.....","...58.3..","2.......5"}
//	args := []string{".....3...", "6....5..8", "7.1.4.2..", "5......4.","..9...1..", ".7...69..", ".9.2.....", "3.......9", ".4.1...75"}
//	args := []string{"....239..", "5.27.....","7.3.6..8.","34...7..8", "..54813..", "68.5...14","......893",".28...15.","1........"}
	if Input(args) {
		sudoku := [9][9]rune{}
//		revsudoku := [9][9]rune{}
		sudoku = Fill(sudoku, args)

//		if Solve(&sudoku) {
//			revsudoku = Fill(revsudoku, args)
//			RevSolve(&revsudoku)
//		}

//		if Solve(&sudoku) && sudoku == revsudoku {
		if Solve(&sudoku) {
			for y := 0; y < 9; y++ {
				for x := 0; x < 9; x++ {
					if x != 8 {
						z01.PrintRune(rune(sudoku[y][x]))
						z01.PrintRune(' ')
					} else {
						z01.PrintRune(rune(sudoku[y][x]))
					}
				}
				z01.PrintRune('\n')
			}
		} else {
			fmt.Println("Error")
		}
	}
}

// Input function confirms if the input is valid,

func Input(args []string) bool {
	if len(args) != 9 {
		fmt.Println("Error")
		return false
	}

	for i := 0; i < len(args); i++ {
		if len(args[i]) != 9 {
			fmt.Println("Error")
			return false
		}
	}
	for i := 0; i < len(args); i++ {
		for _, k := range args[i] {
			if k == 47 || k == 48 {
				fmt.Println("Error")
				return false
			} else if k < 46 || k > 57 {
				fmt.Println("Error")
				return false
			}
		}
	}
	return true
}

// Fill function fills the sudoku array of runes
// with values from the input string array.
func Fill(sudoku [9][9]rune, args []string) [9][9]rune {
	for y := range args {
		for x := range args[y] {
			sudoku[y][x] = rune(args[y][x])
		}
	}
	return sudoku
}

// CountDots function returns true if there are still
// empty cells, and returns false if there are none.
func CountDots(sudoku *[9][9]rune) bool {
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if sudoku[y][x] == '.' {
				return true
			}
		}
	}
	return false
}

// IsValid function iterates over numbers
// that can be potentially placed in empty cells
// and checks if they are valid candidates
// (satisfy each of 3 conditions
// for number placement in sudoku game)
func IsValid(sudoku *[9][9]rune, x int, y int, i rune) bool {
	for j := 0; j < 9; j++ {
		if i == sudoku[j][x] {
			return false
		}
	}

	for k := 0; k < 9; k++ {
		if i == sudoku[y][k] {
			return false
		}
	}

	//square check
	a := x / 3
	b := y / 3

	for l := 3 * a; l < 3*(a+1); l++ {
		for m := 3 * b; m < 3*(b+1); m++ {
			if i == sudoku[m][l] {
				return false
			}
		}
	}
	return true
}

// Solve function is a backtracking algorithm
// that uses recursion to check potential solutions
// starting from 1.
func Solve(sudoku *[9][9]rune) bool {
	if !CountDots(sudoku) {
		return true
	}
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if sudoku[y][x] == '.' {
				for i := '1'; i <= '9'; i++ {
					if IsValid(sudoku, x, y, i) {
						sudoku[y][x] = i
						if Solve(sudoku) {
							return true
						}
					}
					sudoku[y][x] = '.'
				}
				return false
			}
		}
	}
	return false // ? lishnii?
}

// RevSolve function is a backtracking algorithm
// that uses recursion to check potential solutions
// starting from 9.

/*
func RevSolve(revsudoku *[9][9]rune) bool {
	if !CountDots(revsudoku) {
		return true
	}
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if revsudoku[y][x] == '.' {
				for i := '9'; i >= '1'; i-- {
					if IsValid(revsudoku, x, y, i) {
						revsudoku[y][x] = i
						if RevSolve(revsudoku) {
							return true
						}
					}
					revsudoku[y][x] = '.'
				}
				return false
			}
		}
	}
	return false
}
*/
