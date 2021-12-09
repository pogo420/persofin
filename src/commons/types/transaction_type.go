package types

type _transaction_type string

func (tt _transaction_type) Value() string {
	return string(tt)
}

const (
	ToAccount       _transaction_type = "ToAccount"
	FromAccount     _transaction_type = "FromAccount"
	AccountMovement _transaction_type = "AccountMovement"
)

type TransactionType interface {
	Value() string
}
