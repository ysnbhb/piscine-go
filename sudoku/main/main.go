package main

import (
	"fmt"
	"os"
	"strings"
	"piscine"
)

func main() {
	args := os.Args[1:]
	joinedargs := strings.Join(args, " ")
	args = strings.Split(joinedargs, " ")
	if len(args) != 9 {
		fmt.Println("Please provide a 9x9 Sudoku puzzle")
		return
	}
	var tr [][]int
	for i := 0; i < len(args); i++ {
		if len(args[i]) != 9 {
			fmt.Println("error")
			return
		}
		tr = append(tr, piscine.Tonumbre(args[i]))
	}
	for i := 0; i < len(args); i++ {
		for j := 0; j < 9; j++ {
			if tr[i][j] == 0 {
				continue
			}
			if !piscine.Chec(tr, i, j, tr[i][j]) {
				return
			}
		}
	}
	piscine.Print_sudoku(tr)
	fmt.Println("__")
	if !piscine.Solve(tr, 0, 0) {
		fmt.Println("Error")
		return
	}
	piscine.Print_sudoku(tr)
}
