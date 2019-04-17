package chaining

func Go(f func(*Chain)) (c *Chain) {
	c = &Chain{}
	f(c)
	return c
}

func (c *Chain) Reject(err error) *Chain {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.err = err
	c.isDone = true
	if c.catch != nil {
		nc := &Chain{err: c.err}
		c.catch(nc)

		return nc
	}

	return c
}

func (c *Chain) Resolve(val interface{}) *Chain {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.val = val
	c.isDone = true
	if c.then != nil {
		nc := &Chain{val: c.val}
		go c.then(nc)

		return nc
	}

	return c
}

func (c *Chain) Then(cb func(*Chain)) *Chain {
	c.lock.Lock()
	defer c.lock.Unlock()

	if !c.isDone {
		c.then = cb
	} else if c.err == nil {
		c = &Chain{val: c.GetVal()}
		cb(c)
	}

	return c
}

func (c *Chain) Catch(cb func(*Chain)) *Chain {
	c.lock.Lock()
	defer c.lock.Unlock()

	if !c.isDone {
		c.catch = cb
	} else if c.err != nil {
		c = &Chain{err: c.GetError()}
		cb(c)
	}

	return c
}
