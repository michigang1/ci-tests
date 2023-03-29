package lab2

import (
	"fmt"
	"strings"
)

var precedence = map[string]uint8{
	plus:     1,
	minus:    1,
	multiply: 2,
	divide:   2,
	power:    3,
}

// PostfixToInfix converts a postfix expression to an infix expression.
func PostfixToInfix(input string) (output string, err error) {
	v := validator{ValidOperatorExp: `[-\+\*\^\/]`, ValidOperandExp: `(\d+|(\d+[,\.]\d+))`}
	if !v.ValidInput(input) {
		err = fmt.Errorf("invalid input expression")
		return
	}
	var stack []string
	var infixArgs []string
	inputArgs := strings.Split(input, " ")
	if errArgs := v.CheckArgsAmount(inputArgs); errArgs != nil {
		err = errArgs
		return
	}
	for _, arg := range inputArgs {
		if !v.IsOperator(arg) {
			infixArgs = append(infixArgs, arg)
			continue
		}
		operator := arg
		stack = append(stack, operator)
		slicedEnd := len(infixArgs) - 2
		sliced := infixArgs[slicedEnd:]
		infixArgs = infixArgs[:slicedEnd]
		operand1, operand2 := sliced[0], sliced[1]
		// expression has more than one calculating operation
		if len(stack) > 1 {
			prevIndex := len(stack) - 2
			prevValue := stack[prevIndex]

			isPowerOperators := precedence[operator] == 3 && precedence[prevValue] == 3
			higherPriority := precedence[operator] > precedence[prevValue]
			if higherPriority || isPowerOperators {
				if v.IncludesOperator(operand2) {
					operand2 = "(" + operand2 + ")"
				} else {
					operand1 = "(" + operand1 + ")"
				}
			}
		}
		operand := fmt.Sprintf("%s %s %s", operand1, operator, operand2)
		infixArgs = append(infixArgs, operand)
	}
	output = infixArgs[0]
	return
}
