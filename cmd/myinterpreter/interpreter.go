package main

import (
	"fmt"
	"strconv"
)

type Interpreter struct {
}

func (interpreter *Interpreter) interpret(expression Expr) {
	val := interpreter.visitExpr(expression)
	fmt.Println(stringify(val))
}

func stringify(obj any) string {
	if obj == nil {
		return "nil"
	}

	if objStr, ok := obj.(string); ok {
		if val, err := strconv.ParseFloat(objStr, 64); err == nil {
			text := strconv.FormatFloat(val, 'f', -1, 64)
			if lastTwo := text[len(text)-2:]; lastTwo == ".0" {
				return text[0 : len(text)-2]
			}
			return text
		}
	}

	return fmt.Sprintf("%v", obj)
}

func (interpreter *Interpreter) visitExpr(expr Expr) any {
	switch expr.exprType {
	case LITERAL:
		return expr.value
	case GROUPING:
		return interpreter.evaluate(expr.left)
	case UNARY:
		right := interpreter.evaluate(expr.right)
		switch expr.operator.Type {
		case BANG:
			return !isTruthy(right)
		case MINUS:
			checkNumberOperand(expr.operator, right)
			objStr, _ := right.(string)
			val, _ := strconv.ParseFloat(objStr, 64)
			return -val
		}
	case BINARY:
		left := interpreter.evaluate(expr.left)
		right := interpreter.evaluate(expr.right)
		rightStr, _ := right.(string)
		rightVal, _ := strconv.ParseFloat(rightStr, 64)
		leftStr, _ := left.(string)
		leftVal, _ := strconv.ParseFloat(leftStr, 64)
		switch expr.operator.Type {
		// 	case MINUS:
		// 		return leftVal - rightVal
		// 	case PLUS:
		case SLASH:
			return leftVal / rightVal
		case STAR:
			return leftVal / rightVal
			// 	case GREATER:
			// 		return leftVal > rightVal
			// 	case GREATER_EQUAL:
			// 		return leftVal >= rightVal
			// 	case LESS:
			// 		return leftVal < rightVal
			// 	case LESS_EQUAL:
			// 		return leftVal <= rightVal
			// 	case BANG_EQUAL:
			// 		return !isEqual(expr.left, expr.right)
			// 	case EQUAL_EQUAL:
			// 		return isEqual(expr.left, expr.right)
		}
		// 	return nil
	}
	return nil
}

func (interpreter *Interpreter) evaluate(expr *Expr) any {
	return interpreter.visitExpr(*expr)
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
	if objStr, ok := operand.(string); ok {
		if _, err := strconv.ParseFloat(objStr, 64); err == nil {
			return
		}
	}
	panic(runtimeError(operator, "Operand must be a number"))
}

// func checkNumberOperands(operator Token, left any, right any) {
// 	_, okl := left.(float64)
// 	_, okr := right.(float64)

// 	if okl && okr {
// 		return
// 	}

// 	panic(runtimeError(operator, "Operands must be numbers"))
// }

// func isEqual(expr1 *Expr, expr2 *Expr) bool {
// 	return false
// }

type RuntimeError struct {
	token   Token
	message string
}

func runtimeError(token Token, message string) RuntimeError {
	return RuntimeError{token: token, message: message}
}
