package types

type _currency string

func (c _currency) Value() string {
	return string(c)
}

const (
	USD _currency = "USD"
	INR _currency = "INR"
)

type Currency interface {
	Value() string
}
