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
	// test case for comments object validation
	var cmm ty.Comment = "Hello world"
	if cmm != "Hello world" {
		(*t).Fatalf("Issue in setting comment object")
	}
}
