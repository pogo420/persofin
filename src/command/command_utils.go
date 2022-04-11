package command

import (
	cmm "persofin/src/commons"
	"time"
)

// function to validate
func isValidateDate(date_ cmm.Date) bool {
	_, err := time.Parse(cmm.DATE_FORMAT, string(date_))
	if err != nil {
		return false
	} else {
		return true
	}
}
