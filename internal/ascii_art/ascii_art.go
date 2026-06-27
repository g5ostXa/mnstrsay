package ascii_art

import (
	"os"
	"strings"

	"github.com/g5ostXa/mnstrsay/internal/styles"

	"charm.land/lipgloss/v2"
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

func Render() {

	message := "Sweet dreams!"
	if len(os.Args) > 1 {
		message = strings.Join(os.Args[1:], " ")
	}

	lipgloss.Println(styles.MainMsgBoxStyle.Render("", message, ""))
	lipgloss.Println(styles.MainAsciiLogoStyle.Render(mnstrsayArt))
}
