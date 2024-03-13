package piscine

import (
	"fmt"
)

func Chec(board [][]int, row, col, num int) bool {
	blkrow, blkcol := (row/3)*3, (col/3)*3
	for i := row + 1; i < 9; i++ {
		if board[i][col] == num {
			return false
		}
	}
	for i := col + 1; i < 9; i++ {
		if board[row][i] == num {
			return false
		}
	}
	for i := 0; i < 9; i++ {
		if blkcol+i%3 == col && blkrow+i/3 == row {
			continue
		}
		if board[blkrow+i/3][blkcol+i%3] == num {
			return false
		}
	}
	return true
}

func Print_sudoku(bo [][]int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Print(bo[i][j], " ")
		}
		fmt.Println()
	}
}

func Solve(board [][]int, row, col int) bool {
	for i := row; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != 0 {
				continue
			}
			for num := 1; num <= 9; num++ {
				if IsValid(board, i, j, (num)) {
					board[i][j] = (num)
					if Solve(board, i, j+1) {
						return true
					}
					board[i][j] = 0
				}
			}
			return false
		}
	}
	return true
}

func IsValid(board [][]int, row, col int, num int) bool {
	blkrow, blkcol := (row/3)*3, (col/3)*3
	for i := 0; i < 9; i++ {
		if board[i][col] == num || board[row][i] == num ||
			board[blkrow+i/3][blkcol+i%3] == num {
			return false
		}
	}
	return true
}

func Tonumbre(bo string) []int {
	var toint []int

	for i := 0; i < 9; i++ {
		if bo[i] == '.' {
			toint = append(toint, 0)
		} else {
			toint = append(toint, int(bo[i]-'0'))
		}
	}
	return toint
}

func SplitWithSpace(t string) []string {
	tr := []string{}
	o := ""
	for i := 0; i < len(t); i++ {
		if t[i] == ' ' {
			if o != "" {
				tr = append(tr, o)
			}
			o = ""
		} else {
			o += string(t[i])
		}
	}
	if o != "" {
		tr = append(tr, o)
	}
	return tr
}
