package chaining

// Flow chaining funcs
func Flow(fs ...func(*Chain) (interface{}, error)) func(interface{}, error) (c *Chain) {
	return func(src interface{}, err error) (c *Chain) {
		c = New(src, err)
		for _, f := range fs {
			c = c.NextWithFail(f)
		}

		return c
	}
}
