package samplify

import (
	"testing"
)

func TestValidateSurveyURL(t *testing.T) {
	tables := []struct {
		id        string
		inputURL  string
		outputERR error
	}{
		{
			"1",
			"http://www.google.com",
			nil,
		},
		{
			"2",
			"www.google.com",
			nil,
		},
		{
			"3",
			"",
			ErrURLBlank,
		},
		{
			"4",
			"ac",
			ErrURLMinLength,
		},
		{
			"5",
			".google.com",
			ErrURLPrefix,
		},
		{
			"6",
			"http://.google.com",
			ErrURLHostPrefix,
		},
		{
			"7",
			"fhjksdhfjsfhjsf",
			ErrURLHost,
		},
		{
			"8",
			"http://www.abc.com#",
			ErrURLFragment,
		},
		{
			"9",
			"http://www.goo<gle.com",
			ErrURLInvalid,
		},
		{
			"10",
			"http://www.google.com?a=123",
			nil,
		},
	}
	for _, table := range tables {
		e := ValidateSurveyURL(table.inputURL)
		if e != table.outputERR {
			if e == nil {
				t.Errorf("%s validation failed got: nil, want `%s` ", table.id, table.outputERR.Error())
			} else if table.outputERR == nil {
				t.Errorf("%s validation failed got: `%s`, want nil ", table.id, e.Error())
			} else {
				t.Errorf("%s validation failed got: `%s`, want `%s` ", table.id, e.Error(), table.outputERR.Error())
			}
		}
	}

}
