package main

import (
	"fmt"
	"os"

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

	for i := 1; i < 10; i++ {
		for j := 0; j < 9; j++ {
			if os.Args[i][j] == '.' {
				asChanged := TryNumbers(os.Args, i, j)
				if asChanged {
					i = 1
					j = -1
				}
			}
		}
	}

	fmt.Println("\n\nSudoku solved :")
	PrintSudoku(os.Args)
}

func TryNumbers(Args []string, i int, j int) bool {
	column := GetColumn(Args, j)
	line := Args[i]
	square := GetSquare(Args, i, j)

	var possibilities []int
	for i := 1; i < 10; i++ {
		if !NumberInLine(column, i) && !NumberInLine(line, i) && !NumberInLine(square, i) {
			possibilities = append(possibilities, i)
		}
	}

	fmt.Println("\n at ", i, j, " found possibilities : ", possibilities)
	if len(possibilities) == 1 {
		AddNumberAtIndex(Args, i, j, possibilities[0])
		return true
	}
	if len(possibilities) == 0 {
		fmt.Println("Sudoku impossible ?")
	}

	return false
}

func NumberInLine(s string, n int) bool {
	for i := 0; i < len(s); i++ {
		if int(s[i]) == n+48 {
			return true
		}
	}
	return false
}

func GetColumn(Args []string, columnIndex int) string {
	var column string
	for i := 1; i < 10; i++ {
		column += string(Args[i][columnIndex])
	}
	return column
}

func GetSquare(Args []string, LineIndex int, columnIndex int) string {
	var square string
	LineIndex -= (LineIndex - 1) % 3
	columnIndex -= columnIndex % 3
	for i := LineIndex; i < LineIndex+3; i++ {
		for j := columnIndex; j < columnIndex+3; j++ {
			square += string(Args[i][j])
		}
	}
	return square
}

func AddNumberAtIndex(Args []string, i int, j int, n int) {
	var newString string
	for k := 0; k < 9; k++ {
		if k != j {
			newString += string(Args[i][k])
		} else {
			newString += string(n + 48)
		}
	}
	fmt.Println("\n\nChanged line from : ")
	PrintStr(Args[i])
	fmt.Println("\nto ")
	PrintStr(newString)

	Args[i] = newString
}

func PrintSudoku(Args []string) {
	for i := 1; i < 10; i++ {
		PrintStr(Args[i])
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
