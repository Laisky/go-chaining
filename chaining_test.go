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
	v := c.GetInt()
	return v + 1, nil
}

func throwError(c *chaining.Chain) (r interface{}, err error) {
	return c.GetInt(), errors.New("error")
}

func fail(err error) {
	fmt.Println("got error %v", err)
}

func TestSimpleChain(t *testing.T) {
	r := chaining.New(rootChainFunc()).
		Next(plus1).
		Fail(fail).
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
		Next(plus1).
		Next(plus1).
		Next(throwError). // will interupt chain
		Next(plus1).
		Fail(fail). // will recover chain
		Next(plus1).
		Next(plus1)
	expectVal := 4
	if r.GetInt() != expectVal {
		t.Errorf("expect %v, got %v", expectVal, r.GetInt())
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
