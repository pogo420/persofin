package types

import "fmt"

type Transaction struct {
	Id               int
	Timestamp        Ts
	TransactionValue TransactionValue
	Comment          Comment
	To               Account
	From             Account
}

func (t *Transaction) String() string {
	return fmt.Sprintf("Transaction{ Id: %d , Timestamp: %s, TransactionValue: %s, Comment: %s, To: %s, From: %s}",
		(*t).Id,
		(*t).Timestamp.String(),
		(*t).TransactionValue.String(),
		(*t).Comment.String(),
		(*t).To.String(),
		(*t).From.String())
}
