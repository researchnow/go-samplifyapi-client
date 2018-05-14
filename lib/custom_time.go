package samplify

import (
	"strings"
	"time"
)

const ctLayout = "2006/01/02 15:04:05"

// CustomTime ...
type CustomTime struct {
	time.Time
}

// UnmarshalJSON ...
func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if len(s) == 0 || s == "null" {
		ct.Time = time.Time{}
		return
	}
	ct.Time, err = time.Parse(ctLayout, s)
	return
}

//todo
// // MarshalJSON ...
// func (ct *CustomTime) MarshalJSON() ([]byte, error) {
// 	if ct.Time.UnixNano() == nilTime {
// 		return []byte("null"), nil
// 	}
// 	return []byte(fmt.Sprintf("\"%s\"", ct.Time.Format(ctLayout))), nil
// }

var nilTime = (time.Time{}).UnixNano()

// IsSet ...
func (ct *CustomTime) IsSet() bool {
	return ct.UnixNano() != nilTime
}
