package chaining

type FlowFuncs []func(*Chain) (interface{}, error)

// Flow chaining funcs
func Flow(fs FlowFuncs) func(interface{}, error) (c *Chain) {
	return func(src interface{}, err error) (c *Chain) {
		c = New(src, err)
		for _, f := range fs {
			c = c.NextWithFail(f)
		}

		return c
	}
}
