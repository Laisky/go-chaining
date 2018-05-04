package chaining

type Chain struct {
	val interface{}
	err error
}

// New is the root of the channing
// New(func(args...))
func New(val interface{}, err error) *Chain {
	return &Chain{
		val: val,
		err: err,
	}
}

// Next pass the result of the upstream to the next func
// if any error occured at any point of the upstream,
// no downstream will be involved, until there is a `Fail` deal with the error
func (c *Chain) Next(f func(c *Chain) (interface{}, error)) *Chain {
	if c.err != nil {
		return c
	}

	if val, err := f(c); err != nil {
		c.err = err
	} else {
		c.val = val
	}

	return c
}

// Similiar to Next,
// the different is NextWithFail will ignore the error that happended in the upstream.
// You should deal with the fail by yourself
func (c *Chain) NextWithFail(f func(c *Chain) (interface{}, error)) *Chain {
	if val, err := f(c); err != nil {
		c.err = err
	} else {
		c.val = val
	}

	return c
}

// Fail deal with the first error that occured at the upstream of the chain
func (c *Chain) Fail(f func(err error)) *Chain {
	if c.err == nil {
		return c
	}

	f(c.err)
	c.err = nil
	return c
}
