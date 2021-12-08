package types

import "fmt"

type Response string

func (c *Response) String() string {
	return string(*c)
}

func (c *Response) Print() {
	fmt.Println(*c)
}
