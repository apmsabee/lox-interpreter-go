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
		//scan tokens
		scanner := newScanner(config.filename)
		var tokens []Token
		for scanner.current <= len(scanner.fileContents) {
			if token, errMsg := scanner.nextToken(); errMsg == "" {
				if token != nil {
					tokens = append(tokens, *token)
					if token.Type == EOF {
						break
					}
				}
			} else {
				fmt.Fprint(os.Stderr, errMsg)
			}
		}

		parser := newParser(tokens)
		res, err := parser.Parse()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Parsing error: %v\n", err)
		}
		fmt.Printf("%s\n", print_ast(res))
		os.Exit(parser.exitCode)
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
