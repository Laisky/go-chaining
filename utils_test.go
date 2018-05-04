package chaining_test

import (
	"testing"

	chaining "github.com/Laisky/go-chaining"
)

func TestFlow(t *testing.T) {
	var (
		expect int
		c      *chaining.Chain
	)

	// case 1
	expect = 4
	c = chaining.Flow(chaining.FlowFuncs{
		plus1,
		plus1,
		plus1,
		plus1,
	})(0, nil)
	if c.GetError() != nil {
		t.Errorf("got error %+v", c.GetError())
	}
	if c.GetInt() != expect {
		t.Errorf("expect %v got %v", expect, c.GetInt())
	}

	// case 2
	expect = 3
	c = chaining.Flow(chaining.FlowFuncs{
		plus1, // +1
		throwError,
		plus1, // +1
		plus1, // +1
	})(0, nil)
	if c.GetError() == nil {
		t.Error("error occured in the upstream disapears")
	}
	if c.GetError().Error() != "error occured in the upstream" {
		t.Errorf("got error %+v", c.GetError())
	}
	if c.GetInt() != expect {
		t.Errorf("expect %v got %v", expect, c.GetInt())
	}
}
