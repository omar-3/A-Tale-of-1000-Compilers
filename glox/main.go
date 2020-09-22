package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"glox/lox"
)

func main() {
	switch len(os.Args) - 1 {
	case 0:
		runPrompt()
	case 1:
		runFile(os.Args[1])
	default:
		fmt.Fprintln(os.Stderr, "Usage: golox [script]")
		os.Exit(64)
	}
}

func runFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "The file isn't accessible")
		os.Exit(-1)
	}
	defer file.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(file)

	code := buf.String()
	run(code)
	if lox.HadError {
		os.Exit(65)
	}
}

func runPrompt() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		scanner.Scan()
		line := scanner.Text()
		if line == "" {
			break
		}
		run(line)
		lox.HadError = false	
	}
}

func run(source string) {
	scanner := lox.NewScanner(source)
	tokens := scanner.scanTokens()

	for _, token := tokens {
		fmt.Printf("%#v\n", token)
	}
}


