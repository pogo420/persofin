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
	// function for string representation of transaction class

	var to_str string
	var from_str string

	if (*t).To == nil && (*t).From == nil {

		return ""

	} else if (*t).To == nil {

		to_str = "To NA"
		from_str = (*t).From.String()

	} else if (*t).From == nil {

		to_str = (*t).To.String()
		from_str = "From NA"
	} else {
		to_str = (*t).To.String()
		from_str = (*t).From.String()
	}

	return fmt.Sprintf("Transaction{ Id: %d , Timestamp: %s, TransactionValue: %s, Comment: %s, To: %s, From: %s}",
		(*t).Id,
		(*t).Timestamp.String(),
		(*t).TransactionValue.String(),
		(*t).Comment.String(),
		to_str,
		from_str)
}
