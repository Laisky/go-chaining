package chaining_test

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/Laisky/go-chaining"
)

func rootChainFunc() (int, error) {
	return 0, nil
}

func plus1(c *chaining.Chain) (interface{}, error) {
	return c.GetInt() + 1, nil
}

func throwError(c *chaining.Chain) (interface{}, error) {
	return nil, errors.New("error occured in the upstream")
}

func handleError(err error) {
	fmt.Printf("got error %+v\n", err)
}

func TestSimpleChain(t *testing.T) {
	r := chaining.New(rootChainFunc()).
		Next(plus1).
		Fail(handleError).
		Next(plus1).
		Next(plus1)
	expectVal := 3
	if r.GetInt() != expectVal {
		t.Errorf("expect %v, got %v", expectVal, r.GetInt())
	}
}
func TestChainWithError(t *testing.T) {
	r := chaining.New(rootChainFunc()).
		Next(plus1).
		Next(plus1).
		Next(throwError). // will interupt chain
		Next(plus1).
		Next(plus1)
	expectVal := 2
	if r.GetInt() != expectVal {
		t.Errorf("expect %v, got %v", expectVal, r.GetInt())
	}
}
func TestChainWithErrorAndFail(t *testing.T) {
	r := chaining.New(rootChainFunc()).
		Next(plus1).      //+1
		Next(plus1).      // +1
		Next(throwError). // will interupt chain
		Next(plus1).
		Fail(handleError). // will recover chain
		Next(plus1).       // +1
		Next(throwError).
		Next(plus1).
		Next(plus1).
		NextWithFail(plus1) // +1
	expectVal := 4
	if r.GetInt() != expectVal {
		t.Errorf("expect %v, got %v", expectVal, r.GetInt())
	}
	if r.GetError() != nil {
		t.Error("`NextWithFail(plus1)` not vanish error")
	}
}

func toLower(c *chaining.Chain) (interface{}, error) {
	v := c.GetString()
	return strings.ToLower(v), nil
}

func toUpper(c *chaining.Chain) (interface{}, error) {
	v := c.GetString()
	return strings.ToUpper(v), nil
}

func TestChainWithString(t *testing.T) {
	f := func() (string, error) { return "aBcD", nil }
	r := chaining.New(f()).
		Next(toLower).
		Next(toUpper)

	expectVal := "ABCD"
	if r.GetString() != expectVal {
		t.Errorf("expect %v, got %v", expectVal, r.GetString())
	}
}
