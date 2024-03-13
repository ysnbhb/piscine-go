package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	var arr string
	if len(os.Args) == 1 {
		var stdInInput [5000]byte
		n, _ := os.Stdin.Read(stdInInput[:])
		arr = string(stdInInput[:n])
	} else if len(os.Args) > 1 {
		arr = os.Args[1]
	}
	if arr[len(arr)-1] != '\n' {
		arr += "\n"
	}
	var y, x int
	for i := 0; arr[i] != '\n'; i++ {
		x++
	}
	for _, r := range arr {
		if r == '\n' {
			y++
		}
	}
	tr := []string{"quadA", "quadB", "quadC", "quadD", "quadE"}
	if arr[0] == 'o' || arr[0] == '/' || arr[0] == 'A' {
		var t []string
		for i := 0; i < 5; i++ {
			cmd := exec.Command("./"+tr[i], string(x+'0'), string(y+'0'))
			d, _ := cmd.Output()
			if string(arr) == string(d) {
				t = append(t, tr[i])
			}
		}
		if len(t) == 0 {
			fmt.Print("Not a quad function")
		} else {
			for i := 0; i < len(t); i++ {
				fmt.Print("[" + t[i] + "] [" + string(x+'0') + "] [" + string(y+'0') + "]")
				if i != len(t)-1 {
					fmt.Print(" || ")
				}
			}
			fmt.Println()
		}
	} else {
		fmt.Print("Not a quad function")
	}
}
