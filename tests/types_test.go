package tests

import (
	ty "persofin/src/commons/types"
	"testing"
)

func TestTimeStamp(t *testing.T) {
	// test case for timstamp validation
	var ts ty.Ts = "11 Dec 2021 12:42:00"
	if ts.Validate() {
		(*t).Fatalf("Issue in parsing timestamp")
	}
}

func TestTimeStampIssue(t *testing.T) {
	// -ve test case for time stamp
	var ts ty.Ts = "11-Dec-2021 12:42:00"
	if !ts.Validate() {
		(*t).Fatalf("Issue in parsing timestamp")
	}
}

func TestTimeStampValue(t *testing.T) {
	// test case for time stamp value
	var ts ty.Ts = "11 Dec 2021 12:42:00"
	if ts.GetTimeStamp() != "11 Dec 2021 12:42:00" {
		(*t).Fatalf("Issue in parsing timestamp")
	}
}

func TestTimeStampDateValue(t *testing.T) {
	// test case to validate extract date logic of time stamp
	var ts ty.Ts = "11 Dec 2021 12:42:00"
	if ts.GetDate() != "2021-12-11" {
		(*t).Fatalf("Issue in parsing timestamp's date part")
	}
}

func TestTimeStampDateValueIssue(t *testing.T) {
	// -ve test case to validate extract date logic of time stamp
	var ts ty.Ts = "11 Dec 21 12:42:00"
	if ts.GetDate() != "" {
		(*t).Fatalf("Issue in parsing timestamp's date part")
	}
}

func TestComments(t *testing.T) {
	// test case for Comments object validation
	var cmm ty.Comment = "Hello world"
	if cmm != "Hello world" {
		(*t).Fatalf("Issue in setting comment object")
	}
}

func TestCommand(t *testing.T) {
	// test case for Command object validation
	var cmm ty.Command = "Hello world"
	if cmm != "Hello world" {
		(*t).Fatalf("Issue in setting command object")
	}
}

func TestResponse(t *testing.T) {
	// test case for Response object validation
	var cmm ty.Response = "Hello world"
	if cmm != "Hello world" {
		(*t).Fatalf("Issue in setting response object")
	}
}

func TestCurrency(t *testing.T) {
	var curr ty.Currency = ty.USD
	if curr != ty.USD {
		(*t).Fatalf("Issue in setting currency object")
	}
}

func TestAccountValue(t *testing.T) {
	av := &ty.AccountValue{Value: 23, Currency: ty.USD}

	if (*av).String() != "AccountValue { value: 23, currency: USD}" {
		(*t).Fatalf("Issue in setting AccountValue object %s", (*av).String())
	}
}

func TestTransactionValue(t *testing.T) {
	tv := &ty.TransactionValue{Value: 23, Currency: ty.USD}

	if (*tv).String() != "TransactionValue { value: 23, currency: USD}" {
		(*t).Fatalf("Issue in setting TransactionValue object %s", (*tv).String())
	}
}

func TestAccount(t *testing.T) {
	av := &ty.AccountValue{Value: 23, Currency: ty.USD}
	a := &ty.Account{Value: *av, Name: "Ola"}

	if (*a).String() != "Account { name: Ola, currency: AccountValue { value: 23, currency: USD}}" {
		(*t).Fatalf("Issue in setting Account object %s", (*a).String())
	}
}
