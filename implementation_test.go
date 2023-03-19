package lab2

import (
	"fmt"
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestPostfixToInfix(c *C) {
	res, err := PostfixToInfix("2 2 +")
	c.Assert(res, Equals, "2 + 2")

	res, err = PostfixToInfix("1 2 + 4 * 3 +")
	c.Assert(res, Equals, "(1 + 2) * 4 + 3")

	res, err = PostfixToInfix("9 6 3 / - 4 2 - 5 * + 8 4 / + 10 2 + 2 ^ 3 / - 2 * -")
	c.Assert(res, Equals, "9 - 6 / 3 + (4 - 2) * 5 + 8 / 4 - ( - (10 + 2) ^ 2 / 3) * 2")

	res, err = PostfixToInfix(" '''pi3.14''' 2 +/")
	c.Assert(err, ErrorMatches, "invalid input expression")
}

func ExamplePostfixToInfix() {
	res, err := PostfixToInfix("2 2 +")
	if err != nil {
		fmt.Println(err)
	} else {
		panic(err)
	}
	fmt.Println(res)

	// Output:
	// 2 + 2
}
