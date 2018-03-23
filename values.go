package chaining

// GetError get the error
func (c *Chain) GetError() error {
	return c.err
}

// GetVal get the value in interface{}
func (c *Chain) GetVal() interface{} {
	return c.val
}

// GetInt get the value in int
func (c *Chain) GetInt() int {
	return c.val.(int)
}

// GetInt32 get the value in int32
func (c *Chain) GetInt32() int32 {
	return c.val.(int32)
}

// GetInt64 get the value in int64
func (c *Chain) GetInt64() int64 {
	return c.val.(int64)
}

// GetFloat32 get the value in float32
func (c *Chain) GetFloat32() float32 {
	return c.val.(float32)
}

// GetFloat64 get the value in float64
func (c *Chain) GetFloat64() float64 {
	return c.val.(float64)
}

// GetBool get the value in bool
func (c *Chain) GetBool() bool {
	return c.val.(bool)
}

// GetSliceString get the value in []string
func (c *Chain) GetSliceString() []string {
	return c.val.([]string)
}

// GetSliceInterface get the value in []interface{}
func (c *Chain) GetSliceInterface() []interface{} {
	return c.val.([]interface{})
}

// GetMapStringString get the value in map[string]string
func (c *Chain) GetMapStringString() map[string]string {
	return c.val.(map[string]string)
}

// GetMapStringInterface get the value in map[string]interface{}
func (c *Chain) GetMapStringInterface() map[string]interface{} {
	return c.val.(map[string]interface{})
}
