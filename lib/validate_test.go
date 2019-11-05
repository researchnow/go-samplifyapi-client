package samplify_test

import (
	"testing"

	samplify "github.com/researchnow/go-samplifyapi-client/lib"
)

func TestValidateSurveyLink(t *testing.T) {
	tables := []struct {
		id        string
		inputURL  string
		outputERR error
	}{
		{
			"Case 1 : Happy case",
			"http://www.google.com",
			nil,
		},
		{
			"Case 2 : Happy case",
			"www.google.com",
			nil,
		},
		{
			"Case 3 : Balnk URL",
			"",
			samplify.ErrURLBlank,
		},
		{
			"Case 4 : URL with minimum length",
			"ac",
			samplify.ErrURLMinLength,
		},
		{
			"Case 5 : URL has a prefix `.`",
			".google.com",
			samplify.ErrURLPrefix,
		},
		{
			"Case 6 : URL host has a prefix `.`",
			"http://.google.com",
			samplify.ErrURLHostPrefix,
		},
		{
			"Case 7 : URL host is not present",
			"fhjksdhfjsfhjsf",
			samplify.ErrURLHost,
		},
		// {
		// 	"Case 8 : URL has a fragment `#`",
		// 	"http://www.abc.com#",
		// 	samplify.ErrURLFragment,
		// },
		{
			"Case 9 : URL is inavalid with special charcter `<`",
			"http://www.goo<gle.com",
			samplify.ErrURLInvalid,
		},
		{
			"Case 10 : Valid URL with params",
			"http://www.google.com?a=123",
			nil,
		},
		{
			"Case 11 : URL is just a void space ` `",
			" ",
			samplify.ErrURLMinLength,
		},
		{
			"Case 12 : URL is just a collection of void spaces",
			"       ",
			samplify.ErrURLHost,
		},
	}
	for _, table := range tables {
		e := samplify.ValidateSurveyLink(table.inputURL)
		if e != table.outputERR {
			t.Errorf("%s validation failed got: `%v`, want `%v` ", table.id, e, table.outputERR)
		}
	}

}
