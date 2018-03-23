# go-chaining

Simply functions chaining in Go


## Methods

`.New(val interface{}, err error) *Chain`

create the chaining by called a func that returns a value and an error.

`.Next(f func(c *Chain) (interface{}, error)) *Chain`

pass chaning to next func

`.Fail(f func(err error)) *Chain`

if any error has occured in the upstream, all the downstrem will be ignored until thers is a `Fail` handle the error.


## Example

```go
package main

import (
	"errors"
	"fmt"

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
	return c.GetInt(), errors.New("some error happened")
}

func handleError(err error) {
	fmt.Printf("deal with error: %v\n", err)
}

func main() {
	r := chaining.New(rootChainFunc()).  // create chaining
		Next(plus1). // +1
		Next(plus1). // +1
		Next(throwError). // will interupt chain
		Next(plus1).
		Next(plus1).
		Fail(handleError). // will recover chain
		Next(plus1). // +1
		Next(plus1) // +1

	fmt.Printf("got result: %v\n", r.GetInt())
	// got 4
}

```

There are many convenient methods to get the value from chaining:

```
.GetError() error
.GetVal() interface
.GetInt() int
.GetInt32() int32
.GetInt64() int64
.GetFloat32() float32
.GetFloat64() float64
.GetBool() bool
.GetSliceString() []string
.GetSliceInterface() []interface{}
.GetMapStringString() map[string]string
.GetMapStringInterface() map[string]interface{}
```
