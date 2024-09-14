package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/01-edu/z01"
)

// run command :
// go run . ".96.4...1" "1...6...4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7"

func main() {
	if len(os.Args) != 10 {
		fmt.Println("Error")
		return
	}

	for i := 1; i < 10; i++ {
		if len(os.Args[i]) != 9 {
			fmt.Println("Error")
			return
		}
	}

	fmt.Println("Starting to solve : ")
	PrintSudoku(os.Args)

	sudoku := make([]string, len(os.Args))
	copy(sudoku, os.Args)
	var addRandom = false

	for !IsSolved(sudoku) {
		for i := 1; i < 10; i++ {
			for j := 0; j < 9; j++ {
				if sudoku[i][j] == '.' {
					asChanged := TryNumbers(sudoku, i, j, addRandom)
					if asChanged == 1 {
						i = 1
						j = -1
					} else if asChanged == 2 {
						fmt.Println("\n\nWrong solition :")
						PrintSudoku(sudoku)
						copy(sudoku, os.Args)
						addRandom = true

					}
				}
			}
		}
		addRandom = true
	}

	fmt.Println("\n\nSudoku solved :")
	PrintSudoku(sudoku)
}

func TryNumbers(sudoku []string, i int, j int, random bool) int {
	column := GetColumn(sudoku, j)
	line := sudoku[i]
	square := GetSquare(sudoku, i, j)

	var possibilities []int
	for i := 1; i < 10; i++ {
		if !NumberInLine(column, i) && !NumberInLine(line, i) && !NumberInLine(square, i) {
			possibilities = append(possibilities, i)
		}
	}

	if len(possibilities) == 1 {
		AddNumberAtIndex(sudoku, i, j, possibilities[0])
		return 1
	}
	if len(possibilities) == 0 {
		return 2
	}

	if len(possibilities) == 2 && random {
		rng := rand.New(rand.NewSource(time.Now().UnixNano()))
		if rng.Intn(4) == 1 {
			TheChoosenPossibility := rng.Intn(2)
			AddNumberAtIndex(sudoku, i, j, possibilities[TheChoosenPossibility])
		}
	}

	return 0
}

func NumberInLine(s string, n int) bool {
	for i := 0; i < len(s); i++ {
		if int(s[i]) == n+48 {
			return true
		}
	}
	return false
}

func GetColumn(sudoku []string, columnIndex int) string {
	var column string
	for i := 1; i < 10; i++ {
		column += string(sudoku[i][columnIndex])
	}
	return column
}

func GetSquare(sudoku []string, LineIndex int, columnIndex int) string {
	var square string
	LineIndex -= (LineIndex - 1) % 3
	columnIndex -= columnIndex % 3
	for i := LineIndex; i < LineIndex+3; i++ {
		for j := columnIndex; j < columnIndex+3; j++ {
			square += string(sudoku[i][j])
		}
	}
	return square
}

func AddNumberAtIndex(sudoku []string, i int, j int, n int) {
	var newString string
	for k := 0; k < 9; k++ {
		if k != j {
			newString += string(sudoku[i][k])
		} else {
			newString += string(n + 48)
		}
	}

	sudoku[i] = newString
}

func PrintSudoku(sudoku []string) {
	for i := 1; i < 10; i++ {
		PrintStr(sudoku[i])
		fmt.Println("")
	}
}

func PrintStr(s string) {
	for i, r := range s {
		z01.PrintRune(r)
		if i != 8 {
			z01.PrintRune(' ')
		}
	}
}

func IsSolved(sudoku []string) bool {
	for i := 1; i < 10; i++ {
		for j := 0; j < 9; j++ {
			if sudoku[i][j] == '.' {
				return false
			}
		}
	}
	return true
}
