package lab2

import (
	"bytes"
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type CHandlerSuite struct{}

var _ = Suite(&CHandlerSuite{})

func (s *CHandlerSuite) TestComputeHandler(c *C) {
	b := bytes.NewBuffer(make([]byte, 0))

	handler := ComputeHandler{
		Input:  bytes.NewReader([]byte("2 2 +")),
		Output: b,
	}
	err := handler.Compute()

	c.Assert(err, Equals, nil)
	c.Assert(b.String(), Equals, "2 + 2")
}

func (s *CHandlerSuite) TestComputeHandlerHard(c *C) {
	b := bytes.NewBuffer(make([]byte, 0))

	handler := ComputeHandler{
		Input:  bytes.NewReader([]byte("9 3 - 2 * 1 +")),
		Output: b,
	}

	err := handler.Compute()

	c.Assert(err, Equals, nil)
	c.Assert(b.String(), Equals, "(9 - 3) * 2 + 1")
}

func (s *CHandlerSuite) TestComputeHandlerError(c *C) {
	b := bytes.NewBuffer(make([]byte, 0))

	handler := ComputeHandler{
		Input:  bytes.NewReader([]byte("! @ + 5 -")),
		Output: b,
	}

	err := handler.Compute()
	c.Assert(err, ErrorMatches, "invalid input expression")
}
