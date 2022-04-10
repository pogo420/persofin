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

## status 
* In progress

## test cases run
* `go test tests/*.go -v`
* `go test tests/*.go -v | grep PASS: | wc -l`
