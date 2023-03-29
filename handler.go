package lab2

import (
	"io"
	"io/ioutil"
	"strings"
)

// ComputeHandler should be constructed with input io.Reader and output io.Writer.
// Its Compute() method should read the expression from input and write the computed result to the output.
type ComputeHandler struct {
	// TODO: Add necessary fields.
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	// TODO: Implement.
	// Read the expression from input.
	input, err := ioutil.ReadAll(ch.Input)
	if err != nil {
		return err
	}

	// Trim any null bytes from the input.
	trimmedInput := strings.Trim(string(input), "\x00")

	res, err := PostfixToInfix(trimmedInput)
	if err != nil {
		return err
	}

	// Write the result to the output.
	_, err = ch.Output.Write([]byte(res))
	if err != nil {
		return err
	}

	return nil
}
