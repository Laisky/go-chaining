# go-chaining

Simply chaining functions in Go

## Example

```go
package main

import "github.com/Laisky/go-chaining"

var c = &Chain{}

func rootChainFunc() (int, error) {
	return 0, nil
}

func plus1(c *Chain) (interface{}, error) {
	v := c.GetInt()
	return v + 1, nil
}

func throwError(c *Chain) (r interface{}, err error) {
	return c.GetInt(), errors.New("error")
}

func fail(err error) {
	fmt.Println("got error %v", err)
}

func main() {
    r := c.New(rootChainFunc()).
            Next(plus1).
            Next(plus1).
            Next(throwError). // will interupt chain
            Next(plus1).
            Fail(fail). // will recover chain
            Next(plus1).
            Next(plus1)

    // r = 4
}
```
