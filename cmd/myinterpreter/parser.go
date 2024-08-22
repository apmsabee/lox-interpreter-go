package main

import (
	"errors"
	"fmt"
)

type Ast struct { //abstract syntax tree
	nodes []Node
}

func (a *Ast) String() string {
	var ast string
	for _, node := range a.nodes {
		ast += node.String() + "\n"
	}
	return ast
}

type Node interface {
	fmt.Stringer //... as long as an object implements String(), it can be of type Node
}

type Keyword struct {
	value string
}

func (k *Keyword) String() string {
	return k.value
}

type Binary struct {
	left     Node
	right    Node
	operator string
}

func (b *Binary) String() string {
	return "(" + b.operator + " " + b.left.String() + " " + b.right.String() + ")"
}

func parse(source string) (error, *Ast) {
	ast := &Ast{}
	scan := newScanner(source) //tokenize our input
	for scan.current <= len(scan.fileContents) {
		if token, errMsg := scan.nextToken(); errMsg == "" {
			switch token.Type {
			case EOF:
				return nil, ast //finish reading the input, return our ast
			default:
				ast.nodes = append(ast.nodes, &Keyword{value: token.lexeme}) //append the Lexeme as a Node to the list of nodes in the interface
			}
		} else {
			return errors.New("Unexpected character"), nil
		}
	}
	return errors.New("unreachable"), nil
}
