package types

import "fmt"

type Command string

func (c *Command) String() string {
	return string(*c)
}

func (c *Command) Print() {
	fmt.Println(*c)
}
