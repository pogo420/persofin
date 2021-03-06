// File containing static data
package commons

// promt string
const PROMT = "persofin>"

// response constant
const (
	SUCCESS = 0
	FAILURE = -1
)

// command const string
const (
	EXIT_COMMAND                = "bye"
	CREATE_ACCOUNT_COMMAND      = "createacc"
	RENAME_ACCOUNT_COMMAND      = "renameacc"
	ADD_TRANSACTION_COMMAND     = "addtrans"
	ACCOUNT_TRANSACTION_COMMAND = "acctrans"
	GET_BALANCE_COMMAND         = "getbal"
	GET_ALL_BALANCE_COMMAND     = "getallbal"
)

const DATE_FORMAT = "2006-01-02" // date format

const DEFAULT_LOG_FILE = "psf.log" // persofin default log file

// env variables
const TEST_MODE_ENV = "PSF_TEST_MODE" // for persofin test mode
const TEST_MODE_ENABLED = "1"

const ENABLE_LOGGING_ENV = "PSF_LOGGING_FLAG" // for persofin logging mode
const LOGGING_ENABLED = "1"

const LOG_FILE_ENV = "PSF_LOG_FILE" // persofin log file
