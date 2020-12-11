package debug

import (
	"bufio"
	"os"
)

// WaitForKeypress a simple debugger tool. It waits for a keypress.
func WaitForKeypress() {
	_, _ = bufio.NewReader(os.Stdin).ReadBytes('\n')
}
