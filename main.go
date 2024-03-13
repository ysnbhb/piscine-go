package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	var stdInInput [5000]byte
	n, _ := os.Stdin.Read(stdInInput[:])
	str := string(stdInInput[:n])
	if str == "Error" {
		fmt.Println(str)
		return
	}
	cmd := exec.Command("./quadchecker", str)
	d, err := cmd.Output()
	if err != nil {
		fmt.Println("Error quadchecker:", err)
	} else {
		fmt.Print(string(d))
	}
}
