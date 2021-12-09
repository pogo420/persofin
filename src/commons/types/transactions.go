package types

import "fmt"

type Transaction struct {
	Id               int
	Timestamp        Ts
	TransactionValue TransactionValue
	Comment          Comment
	To               *Account
	From             *Account
}

func (t *Transaction) String() string {

	if (*t).To == nil && (*t).From == nil {

		return ""

	} else if (*t).To == nil {

		return fmt.Sprintf("Transaction{ Id: %d , Timestamp: %s, TransactionValue: %s, Comment: %s, To: %s, From: %s}",
			(*t).Id,
			(*t).Timestamp.String(),
			(*t).TransactionValue.String(),
			(*t).Comment.String(),
			"To NA",
			(*t).From.String())

	} else if (*t).From == nil {

		return fmt.Sprintf("Transaction{ Id: %d , Timestamp: %s, TransactionValue: %s, Comment: %s, To: %s, From: %s}",
			(*t).Id,
			(*t).Timestamp.String(),
			(*t).TransactionValue.String(),
			(*t).Comment.String(),
			(*t).To.String(),
			"From NA")
	}

	return fmt.Sprintf("Transaction{ Id: %d , Timestamp: %s, TransactionValue: %s, Comment: %s, To: %s, From: %s}",
		(*t).Id,
		(*t).Timestamp.String(),
		(*t).TransactionValue.String(),
		(*t).Comment.String(),
		(*t).To.String(),
		(*t).From.String())
}
