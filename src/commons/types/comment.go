package types

import "fmt"

type Comment string

func (c *Comment) String() string {
	return string(*c)
}

func (c *Comment) Print() {
	fmt.Println(*c)
}
