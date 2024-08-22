package main

import "fmt"

type ExprType int

const (
	BINARY ExprType = iota
	GROUPING
	UNARY
	LITERAL
)

type Expr struct {
	exprType ExprType
	left     *Expr
	right    *Expr
	operator Token
	value    interface{}
}

func print_ast(expr *Expr) string {
	if expr == nil {
		return ""
	}
	switch expr.exprType {
	case BINARY:
		return parenthesize(expr.operator.lexeme, expr.left, expr.right)

	case GROUPING:
		return parenthesize("group", expr.left)

	case UNARY:
		return parenthesize(expr.operator.lexeme, expr.left)

	case LITERAL:
		if expr.value == nil {
			return "nil"
		}
		return fmt.Sprintf("%v", expr.value)
	}
	return ""
}

func parenthesize(name string, exprs ...*Expr) string {
	res := "(" + name
	for _, expr := range exprs {
		res += " " + print_ast(expr)
	}
	res += ")"
	return res
}
