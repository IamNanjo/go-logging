package ansi

import (
	"fmt"
	"os"

	"golang.org/x/sys/windows"
)

func init() {
	stdout := windows.Handle(os.Stdout.Fd())
	stderr := windows.Handle(os.Stderr.Fd())

	var originalOutMode uint32
	if windows.GetConsoleMode(stdout, &originalOutMode) != nil {
		return
	}

	var originalErrMode uint32
	if windows.GetConsoleMode(stderr, &originalErrMode) != nil {
		return
	}

	newOutMode := originalOutMode | windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING
	newErrMode := originalErrMode | windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING

	if windows.SetConsoleMode(stdout, newOutMode) != nil {
		fmt.Fprintf(os.Stderr, "Failed to enable virtual terminal processing for stdout. Some log output might not look right\n")
		return
	}
	if windows.SetConsoleMode(stderr, newErrMode) != nil {
		fmt.Fprintf(os.Stderr, "Failed to enable virtual terminal processing for stderr. Some log output might not look right\n")
		return
	}
}
