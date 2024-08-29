package main

import (
	"fmt"
	"os"
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
		val := interpreter.evaluate(expr.left)
		//fmt.Fprintf(os.Stderr, "Grouping result: %v, Type: %T\n", val, val)
		return val
	case UNARY:
		right := interpreter.evaluate(expr.right)
		switch expr.operator.Type {
		case BANG:
			return !isTruthy(right)
		case MINUS:
			//checkNumberOperand(expr.operator, right)
			ok, float := isFloatVal(right)
			if ok {
				//fmt.Fprintf(os.Stderr, "Returning float val from unary: %v\n", float)
				return -float
			}
			objStr, _ := right.(string)
			val, _ := strconv.ParseFloat(objStr, 64)
			//fmt.Fprintf(os.Stderr, "Type: %T, value: %v\n", val, val)
			return -val
		}
	case BINARY:
		//this is a bit of a mess, but it seems like number literals are stored as strings,
		//but the nested results of arithmetic operations are stored as floats,
		//so theres some weird conditional conversions we have to do for this stage
		left := interpreter.evaluate(expr.left)
		right := interpreter.evaluate(expr.right)

		rightStr, _ := right.(string)
		leftStr, _ := left.(string)

		var leftVal, rightVal float64

		switch left.(type) {
		case float64:
			leftVal, _ = left.(float64)
		default:
			leftVal, _ = strconv.ParseFloat(leftStr, 64)
		}

		switch right.(type) {
		case float64:
			rightVal, _ = right.(float64)
		default:
			rightVal, _ = strconv.ParseFloat(rightStr, 64)
		}

		switch expr.operator.Type {
		case MINUS:
			//printBinary(leftVal, rightVal, "-")
			//fmt.Fprintf(os.Stderr, "Sub result: %v\n", leftVal-rightVal)
			return leftVal - rightVal
		case PLUS:
			//addition and concatenation need to be dealt with
			//again because of the strangeness of our evaluate() return typing, this is going to be ugly
			//throwing type checking into a helper method
			okR, floatRight := isFloatVal(right)
			okL, floatLeft := isFloatVal(left)
			if okR && okL {
				//printBinary(floatLeft, floatRight, "+")
				//fmt.Fprintf(os.Stderr, "Sum result: %v\n", floatLeft+floatRight)
				return floatLeft + floatRight
			}

			stringRight, okRt := right.(string)
			stringLeft, okLt := left.(string)
			if okRt && okLt {
				return stringLeft + stringRight
			}

		case SLASH:
			return leftVal / rightVal
		case STAR:
			return leftVal * rightVal
		case GREATER:
			return leftVal > rightVal
		case GREATER_EQUAL:
			return leftVal >= rightVal
		case LESS:
			return leftVal < rightVal
		case LESS_EQUAL:
			return leftVal <= rightVal
		case BANG_EQUAL:
			return !interpreter.isEqual(expr.left, expr.right)
		case EQUAL_EQUAL:
			return interpreter.isEqual(expr.left, expr.right)
		}
		// 	return nil
	}
	return nil
}
func printBinary(left any, right any, operator string) {
	fmt.Fprintf(os.Stderr, "%v %s %v\n", left, operator, right)
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

func isFloatVal(val any) (bool, float64) {
	//if the value is either already of type float64, or can be parsed into a float64, return true and the associated value
	if fVal, ok := val.(float64); ok {
		return true, fVal
	}

	if valStr, ok := val.(string); ok {
		if fVal, err := strconv.ParseFloat(valStr, 64); err == nil {
			return true, fVal
		}
	}
	return false, 0
}

// func checkNumberOperand(operator Token, operand any) {
// 	if objStr, ok := operand.(string); ok {
// 		if _, err := strconv.ParseFloat(objStr, 64); err == nil {
// 			return
// 		}
// 	}
// 	panic(runtimeError(operator, "Operand must be a number"))
// }

// func checkNumberOperands(operator Token, left any, right any) {
// 	_, okl := left.(float64)
// 	_, okr := right.(float64)

// 	if okl && okr {
// 		return
// 	}

// 	panic(runtimeError(operator, "Operands must be numbers"))
// }

func (interpreter *Interpreter) isEqual(left *Expr, right *Expr) bool {
	leftType := left.left.operator.Type
	rightType := right.left.operator.Type
	fmt.Fprintf(os.Stderr, "TypeL: %v TypeR: %v\n", leftType, rightType)

	if leftType == rightType {
		//evaluate the expressions
		leftVal := interpreter.evaluate(left)
		rightVal := interpreter.evaluate(right)

		fmt.Fprintf(os.Stderr, "ValL: %v ValR: %v\n", leftVal, rightVal)

		if leftType == STRING {
			lString, _ := leftVal.(string)
			rString, _ := rightVal.(string)
			return lString == rString
		}

		if leftType == NUMBER {
			lFloat, _ := leftVal.(float64)
			rFloat, _ := rightVal.(float64)
			return lFloat == rFloat
		}

		return leftVal == rightVal
	}
	return false

	// if expr1 == nil && expr2 == nil {
	// 	return true
	// }
	// if expr1 == nil {
	// 	return false
	// }
	// //need to do type conversion on these again
	// if reflect.TypeOf(expr1) == reflect.TypeOf(expr2) {
	// 	return expr1 == expr2
	// }
	// return false
}

// type RuntimeError struct {
// 	token   Token
// 	message string
// }

// func runtimeError(token Token, message string) RuntimeError {
// 	return RuntimeError{token: token, message: message}
// }
