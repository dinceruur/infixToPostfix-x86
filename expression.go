package main

import (
	"fmt"
	"io"
	"strings"
)

// IsOperator checks if the given character is operator
func IsOperator(c uint8) bool {
	return strings.ContainsAny(string(c), "+ & - & * & /")
}

// IsOperand checks if the given character is operand
func IsOperand(c uint8) bool {
	return c >= '0' && c <= '9'
}

// GetOperatorWeight returns the weight of the operation.
// The following operations "*", "/" has more weight than the operations "+", "-".
// This operator weight will be used in HasHigherPrecedence(op1 string, op2 string) in order to compare
// operation precedence.
func GetOperatorWeight(op string) int {
	switch op {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	}
	return -1
}

// HasHigherPrecedence compares the given two operations in terms of the their weights.
// Returns true if the first parameter has higher weight more than the second parameter.
func HasHigherPrecedence(op1 string, op2 string) bool {
	op1Weight := GetOperatorWeight(op1)
	op2Weight := GetOperatorWeight(op2)
	return op1Weight >= op2Weight
}

// writeOperation
func writeOperation(s string, w io.Writer) {
	switch s {
	case "*":
		_, _ = w.Write([]byte("pop ax \n"))
		_, _ = w.Write([]byte("pop cx \n"))
		_, _ = w.Write([]byte("mul cx \n"))
		_, _ = w.Write([]byte("push ax \n"))
	case "+":
		_, _ = w.Write([]byte("pop ax \n"))
		_, _ = w.Write([]byte("pop cx \n"))
		_, _ = w.Write([]byte("add ax,cx \n"))
		_, _ = w.Write([]byte("push ax \n"))
	case "-":
		_, _ = w.Write([]byte("pop ax \n"))
		_, _ = w.Write([]byte("pop cx \n"))
		_, _ = w.Write([]byte("sub ax,cx \n"))
		_, _ = w.Write([]byte("push ax \n"))
	case "/":
		_, _ = w.Write([]byte("pop ax \n"))
		_, _ = w.Write([]byte("pop cx \n"))
		_, _ = w.Write([]byte("mul cx \n"))
		_, _ = w.Write([]byte("push ax \n"))
	}
}

// writeOperand
func writeOperand(s string, w io.Writer){
	statement := fmt.Sprintf("push %s \n", s)
	_, _ = w.Write([]byte(statement))
}

// ToPostfix converts the given infix notation to the postfix notation.
func ToPostfix(s string, w io.Writer) string {

	var stack Stack

	postfix := ""

	length := len(s)

	for i := 0; i < length; i++ {

		char := string(s[i])
		// Skip whitespaces
		if char == " " {
			continue
		}

		if char == "(" {
			stack.Push(char)
		} else if char == ")" {
			for !stack.Empty() {
				str := stack.Top()
				if str == "(" {
					break
				}
				postfix += " " + str

				writeOperation(str, w)

				stack.Pop()
			}
			stack.Pop()
		} else if !IsOperator(s[i]) {
			// The character is not an operator, it is
			// operand. Read the following characters until fetch an operator.
			j := i
			number := ""
			for ; j < length && IsOperand(s[j]); j++ {
				number = number + string(s[j])
			}
			postfix += " " + number

			writeOperand(number,w)

			i = j - 1
		} else {
			// If character is operator, pop two elements from stack,
			// perform operation and push the result back.
			for !stack.Empty() {
				top := stack.Top()
				if top == "(" || !HasHigherPrecedence(top, char) {
					break
				}
				postfix += " " + top

				writeOperation(top, w)

				stack.Pop()
			}

			stack.Push(char)
		}
	}

	for !stack.Empty() {
		str := stack.Pop()
		postfix += " " + str

		writeOperation(str, w)
	}

	return strings.TrimSpace(postfix)
}
