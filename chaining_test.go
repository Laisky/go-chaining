package chaining_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/Laisky/go-chaining"
)

var c = &chaining.Chain{}

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
	r := c.New(rootChainFunc()).
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
	r := c.New(rootChainFunc()).
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
	r := c.New(rootChainFunc()).
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
