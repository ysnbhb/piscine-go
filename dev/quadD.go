package main

import (
	"fmt"
	"os"
	"strconv"
)

func QuadD(x, y int) {
	// checking if input is negative
	if x <= 0 || y <= 0 {
		return
	}
	// declaring the string where we gonna draw the squares
	var tab []string
	// i, y = horizontal
	for i := 0; i < y; i++ {
		// init the tab so we can fill each line each before printing
		tab = append(tab, "")
		// i, y = vertical
		for j := 0; j < x; j++ {
			// top left
			if i == 0 && j == 0 {
				tab[i] += "A"
				// bottom left
			} else if j == 0 && i == y-1 {
				tab[i] += "A"
				// bottom right
			} else if i == y-1 && j == x-1 && i != 0 {
				tab[i] += "C"
				// top right
			} else if i == 0 && j == x-1 {
				tab[i] += "C"
				// corners
			} else if i == 0 || i == y-1 || j == 0 || j == x-1 {
				tab[i] += "B"
				// space filler
			} else {
				tab[i] += " "
			}
		}
		// print each line in a newline
		fmt.Println(tab[i])
	}
}

func main() {
	arr := os.Args
	if len(arr)!=3{
		fmt.Print("Error")
		return
	}
	a, _ := strconv.Atoi(arr[1])
	b, _ := strconv.Atoi(arr[2])
	QuadD(a, b)
}
