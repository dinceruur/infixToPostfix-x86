package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	fmt.Print("Enter an infix expression: ")

	// Package bufio implements buffered I/O.
	// NewReader method takes a reader interface as an argument and returns a new buffered reader.
	// This program uses os.Stdin which is "/dev/stdin".
	reader := bufio.NewReader(os.Stdin)

	// ReadString reads until the first occurrence of "\n" in the input.
	s, err := reader.ReadString('\n')

	// Trim spaces in the read input
	s = strings.TrimSpace(s)

	// If an error occurs while reading console arguments, the program logs the error
	// and terminates the program with log.Fatal(). log.Fatal prints its arguments and
	// terminates the program with os.Exit(1).
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Ya you postfix notation:", ToPostfix(s))
	return
}
