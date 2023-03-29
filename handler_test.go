package lab2

import (
	"bytes"
	. "gopkg.in/check.v1"
	"testing"
)

func TestHandler(t *testing.T) { TestingT(t) }

type HandlerSuite struct{}

var _ = Suite(&HandlerSuite{})

func (s *HandlerSuite) TestComputeHandler(c *C) {
	b := bytes.NewBuffer(make([]byte, 0))

	handler := ComputeHandler{
		Input:  bytes.NewReader([]byte("2 2 +")),
		Output: b,
	}
	err := handler.Compute()

	c.Assert(err, Equals, nil)
	c.Assert(b.String(), Equals, "2 + 2")
}

func (s *HandlerSuite) TestComputeHandlerHard(c *C) {
	b := bytes.NewBuffer(make([]byte, 0))

	handler := ComputeHandler{
		Input:  bytes.NewReader([]byte("9 3 - 2 * 1 +")),
		Output: b,
	}

	err := handler.Compute()

	c.Assert(err, Equals, nil)
	c.Assert(b.String(), Equals, "(9 - 3) * 2 + 1")
}

func (s *HandlerSuite) TestComputeHandlerError(c *C) {
	b := bytes.NewBuffer(make([]byte, 0))

	handler := ComputeHandler{
		Input:  bytes.NewReader([]byte("! @ + 5 -")),
		Output: b,
	}

	err := handler.Compute()
	c.Assert(err, ErrorMatches, "invalid input expression")
}
