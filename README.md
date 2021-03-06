# go-chaining

Simply functions chaining in Go


## Methods

`.New(val interface{}, err error) *Chain`

create the chaining by called a func that returns a value and an error.

`.Next(f func(c *Chain) (interface{}, error)) *Chain`

pass chaning to next func, will ignore funcs if there is any error occured on the upstream

`.NextWithFail(f func(c *Chain) (interface{}, error)) *Chain`

pass chaning to next func no matter whether there got an error

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
	return c.GetInt(), nil
}

func throwError(c *chaining.Chain) (interface{}, error) {
	return nil, errors.New("some error happened")
}

func handleError(err error) {
	fmt.Printf("deal with error: %v\n", err)
}

func main() {
	// r := chaining.New(0, nil).  // create chaining manually
	r := chaining.New(rootChainFunc()).  // create chaining by func
		Next(plus1). // +1
		Next(plus1). // +1
		Next(throwError). // will interupt chain
		Next(plus1).
		Next(plus1).
		Fail(handleError). // will recover chain
		Next(plus1). // +1
		Next(throwError).  // interupt again
		Next(plus1).
		NextWithFail(plus1).  // +1, deal error by yourself
		Next(throwError)  // throw error

	r.GetInt() // got 4
	r.GetError()  // got Error("some error happened")
}
```

Or you can use `Flow`:

```go
// Flow chaining your funcs.
// but you should deal with your error by yourself in funcs
Flow(fs ...func(*Chain) (interface{}, error)) func(interface{}, error) (c *Chain)
```

```go
c := chaining.Flow(
	plus1, // +1
	throwError,  // error will pass to the downstream funcs
	plus1, // +1, and vanish the error (you should return error manually in the func)
	plus1, // +1
)(0, nil)

c.GetInt()  // got 4
c.GetError()  // got nil
```


## API References

### `chaining`

```go
type Chain struct

New(val interface{}, err error) *Chain
Flow(fs ...func(*Chain) (interface{}, error)) func(interface{}, error) (c *Chain)
```


### `chaining.Chain`

There are many convenient methods to get the value from chaining:

```go

.Next(f func(c *Chain) (interface{}, error)) *Chain
.NextWithFail(f func(c *Chain) (interface{}, error)) *Chain
.Fail(f func(err error)) *Chain

.GetError() error
.GetVal() interface{}
.GetString() string
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
