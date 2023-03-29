package lab2

import (
	"fmt"
	. "gopkg.in/check.v1"
	"testing"
)

type ParserSuite struct{}

var _ = Suite(&ParserSuite{})

func TestParser(t *testing.T) { TestingT(t) }

func (s *ParserSuite) TestPostfixToInfix(c *C) {
	res, err := PostfixToInfix("2 2 +")
	c.Assert(res, Equals, "2 + 2")

	res, err = PostfixToInfix("1 2 + 4 * 3 +")
	c.Assert(res, Equals, "(1 + 2) * 4 + 3")

	res, err = PostfixToInfix("34 6 9 7 3 10 7 - / ^ + * -")
	c.Assert(res, Equals, "34 - 6 * (9 + 7 ^ (3 / (10 - 7)))")

	res, err = PostfixToInfix(" '''pi3.14''' 2 +/")
	c.Assert(err, ErrorMatches, "invalid input expression")

	res, err = PostfixToInfix("2 + 2 ^ /")
	c.Assert(err, ErrorMatches, "too many operators")

	res, err = PostfixToInfix("2 2 2 +")
	c.Assert(err, ErrorMatches, "too many operands")
}

func ExamplePostfixToInfix() {
	res, err := PostfixToInfix("2 2 +")
	if err != nil {
		panic(err)
	} else {
		fmt.Println(res)
	}
	// Output:
	// 2 + 2
}
