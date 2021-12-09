package types

import "fmt"

type Transaction struct {
	Id               int
	TransactionType  TransactionType
	Timestamp        Ts
	TransactionValue TransactionValue
	Comment          Comment
	To               Account
	From             Account
}

func (t *Transaction) String() string {
	return fmt.Sprintf("Transaction{ Id: %d , TransactionType: %s, Timestamp: %s, TransactionValue: %s, Comment: %s, To: %s, From: %s}",
		(*t).Id,
		(*t).TransactionType.Value(),
		(*t).Timestamp.String(),
		(*t).TransactionValue.String(),
		(*t).Comment.String(),
		(*t).To.String(),
		(*t).From.String())
}
