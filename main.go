package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// handleError logs the error and terminates the program with
// status code 1: log.Fatal terminates the program by executing the os.Exit(1) statement.
func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	fmt.Print("Enter an infix expression: ")

	// Package bufio implements buffered I/O.
	// NewReader method takes a reader interface as an argument and returns a new buffered reader.
	// This program uses os.Stdin which is "/dev/stdin".
	reader := bufio.NewReader(os.Stdin)

	// ReadString reads until the first occurrence of "\n" in the input.
	s, err := reader.ReadString('\n')
	handleError(err)

	// Trim spaces in the read string
	s = strings.TrimSpace(s)

	// Creating the file postfix.asm in the current directory with mode 0666 (default).
	f, err := os.Create("postfix.asm")
	handleError(err)

	// Closing fo after the writing operations ends.
	// defer postpones the function call until the covering function returns.
	defer func() {
		if err := f.Close(); err != nil {
			handleError(err)
		}
	}()

	p := ToPostfix(s, f)

	fmt.Println("Ya you postfix notation:", p)
	return
}
