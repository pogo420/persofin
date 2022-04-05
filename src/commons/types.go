package commons

// data type for transaction amount
type TansactionValue int

// data type for date
type Date string

// Data type for account name(must be unique)
type AccountName string

// Data type for list of accounts
type AccountStats map[AccountName]TansactionValue
