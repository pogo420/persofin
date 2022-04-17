[![Build Status](https://github.com/pogo420/persofin/actions/workflows/persofin_flow.yml/badge.svg?branch=master)](https://github.com/pogo420/persofin/actions/workflows/persofin_flow.yml)

# persofin
Personal finance application - go based application 

## assumptions 
* We can create accounts, they are like small money bucket( eg. income savings , expenditure, etc)
* We can transact amount within the buckets.
* We can consume or add from/to bucket, eg. income or expenditure.
* We can not delete a account once created.
* We can not delete a transaction once created.
* We need to balance a wrong transaction by its complement transaction.
* Account balance can not be negative(checks are there).

## Adding new command  
* Implement interface `BaseCommand`
* Update const group in `static_data.go` with command 
* Update map `Commands` in `command_parser.go`

## Adding new DB/Interface
* Implement `BaseDbInterface`.
* Write your logic. Read comments for error status/return types.
* Update variable `PROD_INTERFACE` with addrerss of new struct(new DB/Iinterface).
* Chech `DummyDbInterface` for ideas(Used for testing).

## test cases run
* `go test tests/*.go -v`
* `go test tests/*.go -v | grep PASS: | wc -l`

## Improvements
* Transaction support, getting all transactions.
