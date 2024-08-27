package main

import (
	"fmt"
	"strconv"
)

type Interpreter struct {
}

func (interpreter *Interpreter) visitExpr(expr Expr) any {
	switch expr.exprType {
	case LITERAL:
		return expr.value
		// case GROUPING:
		// 	return evaluate(expr.left)
		// case UNARY:
		// 	right := evaluate(expr.right)
		// 	switch expr.operator.Type{
		// 	case BANG:
		// 		return !isTruthy(right)
		// 	case MINUS:
		//checkNumberOperand(expr.operator, right)
		// 		return -(float64)(right)
		// 	}
		// case BINARY:
		// 	left := evaluate(expr.left)
		// 	right := evaluate(expr.right)
		// 	switch expr.operator.Type{
		// 	case MINUS:
		// 		return (float64)(left) - (float64)(right)
		// 	case PLUS:

		// 	case SLASH:
		// 		return (float64)(left) / (float64)(right)
		// 	case STAR:
		// 		return (float64)(left) / (float64)(right)
		// 	case GREATER:
		// 		return (float64)(left) > (float64)(right)
		// 	case GREATER_EQUAL:
		// 		return (float64)(left) >= (float64)(right)
		// 	case LESS:
		// 		return (float64)(left) < (float64)(right)
		// 	case LESS_EQUAL:
		// 		return (float64)(left) <= (float64)(right)
		// 	case BANG_EQUAL:
		// 		return !isEqual(expr.left, expr.right)
		// 	case EQUAL_EQUAL:
		// 		return isEqual(expr.left, expr.right)
		// 	}
		// 	return nil
	}
	return nil
}

func isEqual(expr1 *Expr, expr2 *Expr) bool {
	return false
}

func isTruthy(right any) bool {
	if right == nil {
		return false
	}
	val, ok := right.(bool)
	if ok {
		return val
	}
	return true //default return is true
}

func checkNumberOperand(operator Token, operand any) {
	if _, ok := operand.(float64); ok {
		return
	}
	panic(runtimeError(operator, "Operand must be a number"))
}

func checkNumberOperands(operator Token, left any, right any) {
	_, okl := left.(float64)
	_, okr := right.(float64)

	if okl && okr {
		return
	}

	panic(runtimeError(operator, "Operands must be numbers"))
}

type RuntimeError struct {
	token   Token
	message string
}

func runtimeError(token Token, message string) RuntimeError {
	return RuntimeError{token: token, message: message}
}

// func evaluate(expr *Expr) any {
// 	switch expr.exprType {

// 	}
// }

func (interpreter *Interpreter) interpret(expression Expr) {
	val := interpreter.visitExpr(expression)
	fmt.Println(stringify(val))
}

func stringify(obj any) string {
	if obj == nil {
		return "nil"
	}

	if dblObj, ok := obj.(float64); ok {
		text := strconv.FormatFloat(dblObj, 'f', -1, 64)
		fmt.Println(text[len(text)-2:])
		if lastTwo := text[len(text)-2:]; lastTwo == ".0" {
			return text[0 : len(text)-2]
		}
		return text
	}
	return fmt.Sprintf("%v", obj)
}
