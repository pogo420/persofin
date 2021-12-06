package types

import (
	"fmt"
	"time"
)

const TIME_FORMAT = "2 Jan 2006 15:04:05"

type Ts string

func (s *Ts) String() string {
	// stringer method
	return string(*s)
}
func (s *Ts) GetTimeStamp() Ts {
	// getter
	return *s
}
func (s *Ts) GetDate() string {
	// extract date logic
	ts, error := time.Parse(TIME_FORMAT, string(*s))
	if error != nil {
		return ""
	}
	return fmt.Sprintf("%d-%d-%d", ts.Year(), ts.Month(), ts.Day())
}

func (s *Ts) Validate() bool {
	// time stamp logic validation
	_, error := time.Parse(TIME_FORMAT, string(*s))

	return error != nil
}
