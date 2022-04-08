package commons

// data type for transaction amount
type TansactionValue int

// data type for date
type Date string

// Data type for account name(must be unique)
type AccountName string

// Data type for list of accounts
type AccountStats map[AccountName]TansactionValue

// Data type for interface version
type InterfaceVersion string

// TODO: transaction support
// Data type of transaction object
// type Transaction struct {
// 	accountName      AccountName
// 	transactionValue TansactionValue
// 	data             Date
// 	comment          string
// }
