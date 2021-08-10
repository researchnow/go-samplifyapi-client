package samplify

import (
	"encoding/json"
	"testing"
	"time"
)

//marshaling then unmarshaling the same object should work
func TestCustomTime_MarshalUnmarshalJSON(t *testing.T) {
	ct := &CustomTime{
		Time: time.Now(),
	}
	marshaledBytes, err := json.Marshal(ct)
	if err != nil {
		t.Fatalf("error occurred marshaling CustomTime: %s", err)
	}
	newTime := &CustomTime{}
	err = json.Unmarshal(marshaledBytes, newTime)
	if err != nil {
		t.Fatalf("error occurred unmarshaling CustomTime: %s", err)
	}
}
