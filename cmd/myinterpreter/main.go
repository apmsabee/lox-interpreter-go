package main

import (
	"fmt"
	"os"
)

type Config struct {
	command  string
	filename string
}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Fprintln(os.Stderr, "Logs from your program will appear here!")

	config := getConfig()

	scanner := newScanner(config.filename)

	cleanRun := true

	for scanner.current <= len(scanner.fileContents) {
		if token, errMsg := scanner.nextToken(); errMsg == "" {
			fmt.Println(token)
		} else {
			fmt.Fprint(os.Stderr, errMsg)
			cleanRun = false
		}
	}

	if cleanRun {
		os.Exit(0)
	} else {
		os.Exit(65)
	}
}

func getConfig() (config Config) {

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	config.command = os.Args[1]

	if config.command != "tokenize" {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", config.command)
		os.Exit(1)
	}

	config.filename = os.Args[2]
	return

}
