package core

import (
	"fmt"
	"os/exec"
)

func Clear() {

	cmd := exec.Command("clear")

	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("\033[2J\033[H")
	}
	fmt.Println(string(output))
}
