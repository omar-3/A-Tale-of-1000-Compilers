package lox

import (
	"fmt"
	"os"
)

var (
	// HadError is a global variable to be checked
	// after each run of the lexing or parsing
	HadError = false
)

// Error is a wrapper around Report function
// that specifies the line of the error and
// what is the error and where it is on the line
func Error(line int, message string) {
	Report(line, "", message)
}

// Report function is as implied by its name
// is to report errors that happens in the program
func Report(line int, where string, message string) {
	fmt.Fprintf(os.Stderr, "[line %d] Error %s: %s", line, where, message)
	HadError = true
}
