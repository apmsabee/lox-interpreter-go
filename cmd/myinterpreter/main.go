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

	switch config.command {

	case "tokenize":
		scanner := newScanner(config.filename)

		for scanner.current <= len(scanner.fileContents) {
			if token, errMsg := scanner.nextToken(); errMsg == "" {
				if token != nil {
					fmt.Println(token)
					if token.Type == EOF {
						break
					}
				}
			} else {
				fmt.Fprint(os.Stderr, errMsg)
			}
		}

		os.Exit(scanner.exitCode)

	case "parse":
		ok, ast := parse(config.filename)
		if ok != nil {
			fmt.Fprintf(os.Stderr, "Error parsing file: %v\n", ok)
			os.Exit(1)
		}
		fmt.Print(ast.String())

	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", config.command)
		os.Exit(1)
	}

}

func getConfig() (config Config) {

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	config.command = os.Args[1]
	config.filename = os.Args[2]
	return

}
