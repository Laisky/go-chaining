package chaining

type Chain struct {
	val interface{}
	err error
	p   *Chain
}

func (c *Chain) New(val interface{}, err error) *Chain {
	return &Chain{
		val: val,
		err: err,
		p:   c,
	}
}

func (c *Chain) GetVal() interface{} {
	return c.val
}

func (c *Chain) GetInt() int {
	return c.val.(int)
}

func (c *Chain) GetError() error {
	return c.err
}

func (c *Chain) Next(f func(val interface{}) (interface{}, error)) *Chain {
	if c.err != nil {
		return c
	}

	return c.New(f(c.val))
}

func (c *Chain) Fail(f func(err error)) *Chain {
	if c.err == nil {
		return c
	}

	return c
}
