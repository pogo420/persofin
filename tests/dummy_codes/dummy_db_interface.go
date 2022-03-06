package dummy_codes

import ty "persofin/src/commons/types"

// dummy Db interface

type DummyDbInterface struct {
}

func (di *DummyDbInterface) ReadAccountsTable(account_name string) *ty.Account {
	av := &ty.AccountValue{Value: 23, Currency: ty.USD}
	a := &ty.Account{Value: *av, Name: "Ola"}
	return a
}

func (di *DummyDbInterface) WriteAccountsTable(account *ty.Account) int {
	return 1
}

func (di *DummyDbInterface) ReadTransactionTableId(id int) *ty.Transaction {
	trans := &ty.Transaction{Id: 32,
		Timestamp:        "11 Dec 2021 12:42:00",
		TransactionValue: ty.TransactionValue{Value: 23, Currency: ty.USD},
		Comment:          "Hello Transaction",
		To:               &ty.Account{Name: "First", Value: ty.AccountValue{Value: 23, Currency: ty.USD}},
		From:             &ty.Account{Name: "Second", Value: ty.AccountValue{Value: 23, Currency: ty.USD}}}
	return trans
}

func (di *DummyDbInterface) ReadTransactionTableByDate(s ty.Ts, e ty.Ts) *([]ty.Transaction) {
	trans1 := ty.Transaction{Id: 32,
		Timestamp:        "11 Dec 2021 12:42:00",
		TransactionValue: ty.TransactionValue{Value: 23, Currency: ty.USD},
		Comment:          "Hello Transaction",
		To:               &ty.Account{Name: "First", Value: ty.AccountValue{Value: 23, Currency: ty.USD}},
		From:             &ty.Account{Name: "Second", Value: ty.AccountValue{Value: 23, Currency: ty.USD}}}

	trans2 := ty.Transaction{Id: 32,
		Timestamp:        "11 Dec 2021 12:42:00",
		TransactionValue: ty.TransactionValue{Value: 23, Currency: ty.USD},
		Comment:          "Hello Transaction",
		To:               &ty.Account{Name: "First", Value: ty.AccountValue{Value: 23, Currency: ty.USD}},
		From:             nil}

	return &[]ty.Transaction{trans1, trans2}
}

func (di *DummyDbInterface) WriteTransactionTable(*ty.Transaction) int {
	return 1
}
