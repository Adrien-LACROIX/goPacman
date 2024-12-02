package screen

import (
	"fmt"
)

// ClearScreen efface l'écran (peut ne pas fonctionner selon l'OS)
func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
