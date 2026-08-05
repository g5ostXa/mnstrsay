package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"charm.land/lipgloss/v2"
)

var (
	mainTitle = "mnstrsay"
	version   = "dev"
)

// Styles and colors defenitions
var (
	white = lipgloss.Color("#FFFFFF")
	alpha = lipgloss.Color("#7f23db")

	mainMsgBoxStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground((white)).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground((alpha)).
			MarginLeft(2)

	mainAsciiLogoStyle = lipgloss.NewStyle().
				Bold(false).
				Foreground((white))
)

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

	if runtime.GOOS == "linux" {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		_ = cmd.Run()
	}
	fmt.Printf("\033[2J\033[H")
}

func main() {

	showVersion := flag.Bool("version", false, "print current version")

	flag.Parse()
	if *showVersion {
		fmt.Println(mainTitle, strings.TrimSpace(version))
		os.Exit(0)
	}

	clearScreen()
	message := "Sweet dreams!"

	if len(os.Args) > 1 {
		message = strings.Join(os.Args[1:], " ")
	}

	lipgloss.Println(mainMsgBoxStyle.Render("", message, ""))
	lipgloss.Println(mainAsciiLogoStyle.Render(mnstrsayArt))
}
