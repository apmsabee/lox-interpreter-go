package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Fprintln(os.Stderr, "Logs from your program will appear here!")
	l := log.New(os.Stderr, "", 1)

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1]

	if command != "tokenize" {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	// Uncomment this block to pass the first stage

	filename := os.Args[2]
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	tokens := map[rune]string{
		'(': "LEFT_PAREN ( null",
		')': "RIGHT_PAREN ) null",
		'{': "LEFT_BRACE { null",
		'}': "RIGHT_BRACE } null",
		'*': "STAR * null",
		'.': "DOT . null",
		',': "COMMA , null",
		'+': "PLUS + null",
		'-': "MINUS - null",
		';': "SEMICOLON ; null",
	}

	convertedContents := (string)(fileContents)
	cleanRun := true

	if len(convertedContents) > 0 {
		for index, char := range convertedContents {
			if scanned, validFile := tokens[char]; validFile {
				fmt.Println(scanned)
			} else {
				line := strings.Count(convertedContents[0:index], "\n") + 1
				l.Printf("[Line %d] Error: Unexpected character: %c", line, char)
				cleanRun = false
			}

		}
		fmt.Println("EOF  null")
	} else {
		fmt.Println("EOF  null") // Placeholder, remove this line when implementing the scanner
	}

	if cleanRun {
		os.Exit(0)
	} else {
		os.Exit(65)
	}
}
