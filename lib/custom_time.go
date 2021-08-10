package samplify

import (
	"encoding/json"
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

var nilTime = (time.Time{}).UnixNano()

// IsSet ...
func (ct *CustomTime) IsSet() bool {
	return ct.UnixNano() != nilTime
}

func (ct *CustomTime) MarshalJSON() (text []byte, err error) {
	timeString := ct.Time.Format(ctLayout)
	return json.Marshal(timeString)
}