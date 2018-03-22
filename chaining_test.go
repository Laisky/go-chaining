package chaining

import (
	"fmt"
	"testing"
)

var c = &Chain{}

func demo() (int, error) {
	return 1, nil
}

func chan1(val interface{}) (interface{}, error) {
	v := val.(int)
	return v + 1, nil
}

func fail(err error) {
	fmt.Println("got error %v", err)
}

func TestChain(t *testing.T) {
	r := c.New(demo()).Next(chan1).Fail(fail).Next(chan1).Next(chan1)
	expectVal := 4
	if r.GetInt() != expectVal {
		t.Errorf("expect %v, got %v", expectVal, r.GetInt())
	}
}
