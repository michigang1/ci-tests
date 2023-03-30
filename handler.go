package lab2

import (
	"fmt"
	"io"
)

// ComputeHandler should be constructed with input io.Reader and output io.Writer.
// Its Compute() method should read the expression from input and write the computed result to the output.
type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	stack := make([]byte, 0)
	input := make([]byte, 8)
	for {
		n, err := ch.Input.Read(input)
		stack = append(stack, input[:n]...)
		if err == io.EOF {
			break
		}
	}
	result, err := PostfixToInfix(string(stack))
	if err != nil {
		return err
	}
	ch.Output.Write([]byte(result))
	fmt.Println("")
	return nil
}
