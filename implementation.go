package lab2

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	plus     = "+"
	minus    = "-"
	multiply = "*"
	divide   = "/"
	power    = "^"
)

type validator struct {
	ValidOperatorExp string
	ValidOperandExp  string
}

func (v *validator) ValidInput(input string) bool {
	validator := fmt.Sprintf(`^((%s|%s)\s){2,}(%s\s){0,}%s$`, v.ValidOperandExp, v.ValidOperatorExp, v.ValidOperatorExp, v.ValidOperatorExp)
	isValid, _ := regexp.MatchString(validator, input)
	return isValid
}

func (v *validator) CheckArgsAmount(args []string) error {
	operators, operands := 0, 0
	operator := fmt.Sprintf(`^%s$`, v.ValidOperatorExp)
	operand := fmt.Sprintf(`^%s$`, v.ValidOperandExp)
	for _, arg := range args {
		if isOperator, _ := regexp.MatchString(operator, arg); isOperator {
			operators++
		} else if isOperand, _ := regexp.MatchString(operand, arg); isOperand {
			operands++
		}
	}
	switch {
	case operators+operands != len(args):
		return fmt.Errorf("invalid expression argument(s)")
	case operands > operators+1:
		return fmt.Errorf("too many operands")
	case operators > operands-1:
		return fmt.Errorf("too many operators")
	default:
		return nil
	}
}

func (v *validator) IncludesOperator(str string) bool {
	includes, _ := regexp.MatchString(v.ValidOperatorExp, str)
	return includes
}

func (v *validator) IsOperator(str string) bool {
	matching := fmt.Sprintf(`^%s$`, v.ValidOperatorExp)
	includes, _ := regexp.MatchString(matching, str)
	return includes
}

// PostfixToInfix converts a postfix expression to an infix expression.
func PostfixToInfix(input string) (output string, err error) {
	var precedence = map[string]uint8{
		plus:     1,
		minus:    1,
		multiply: 2,
		divide:   2,
		power:    3,
	}

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
