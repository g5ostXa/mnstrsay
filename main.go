package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"charm.land/lipgloss/v2"
)

var white = lipgloss.Color("#FFFFFF")
var alpha = lipgloss.Color("#7f23db")

var mainMsgBoxStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground((white)).
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground((alpha)).
	MarginLeft(2)

var mainAsciiLogoStyle = lipgloss.NewStyle().
	Bold(false).
	Foreground((white))

var mnstrsayArt = `
         \ 
          \
           \ 
       /\_____/\
      |         |
      |  X   X  |
     <     -     >
     (           )
      \/-vvvv-\/
       )      (
       {######}
        \____/
`

func clearScreen() {

	cmd := exec.Command("clear")

	output, err := cmd.Output()
	if err != nil {
		fmt.Print("\033[2J\033[H")
	}
	fmt.Println(string(output))
}

func main() {

	clearScreen()
	message := "Sweet dreams!"

	if len(os.Args) > 1 {
		message = strings.Join(os.Args[1:], " ")
	}

	lipgloss.Println(mainMsgBoxStyle.Render("", message, ""))
	lipgloss.Println(mainAsciiLogoStyle.Render(mnstrsayArt))
}
