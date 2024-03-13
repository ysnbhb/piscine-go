package main

import (
	"fmt"
	"os"
	"piscine"
)

func main() {
	arr := os.Args[1:]
	if len(arr) == 1 {
		arr = piscine.SplitWithSpace(arr[0])
		//use split whit space in contoin if (go run . "".96.4...1" "1...6...4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...0"")
	}
	if len(arr) != 9 {
		fmt.Println("Please provide a 9x9 Sudoku puzzle")
		return
	}
	var tr [][]int
	for i := 0; i < len(arr); i++ {
		if len(arr[i]) != 9 {
			fmt.Println("error")
			return
		}
		tr = append(tr, piscine.Tonumbre(arr[i]))
	}
	for i := 0; i < len(arr); i++ {
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
