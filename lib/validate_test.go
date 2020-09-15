package samplify_test

import (
	"encoding/json"
	"testing"

	samplify "github.com/morningconsult/go-samplifyapi-client/lib"
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

// TestValidateQuotaPlan testing quota plan validation
func TestValidateQuotaPlan(t *testing.T) {
	tables := []struct {
		name     string
		input    string
		expected error
	}{
		{

			"Case 1: Happy path, everything is valid with full quota plan",
			`{ "filters": [{ "attributeId": "13", "options": [ "18-99" ] }], "quotaGroups": [{ "name": "group 1", "quotaCells": [{ "quotaNodes": [{ "attributeId": "11", "options": [ "1" ] }], "perc": 50 }, { "quotaNodes": [{ "attributeId": "11", "options": [ "2" ] }], "perc": 50 } ] }] }`,
			nil,
		},
		{
			"Case 2: Happy path, everything is valid with only filters",
			`{ "filters": [{ "attributeId": "13", "options": [ "18-99" ] }]}`,
			nil,
		},
		{
			"Case 3: Happy path, everything is valid with only quotagroups",
			`{ "quotaGroups": [{ "name": "group 1", "quotaCells": [{ "quotaNodes": [{ "attributeId": "11", "options": [ "1" ] }], "perc": 50 }, { "quotaNodes": [{ "attributeId": "11", "options": [ "2" ] }], "perc": 50 } ] }] }`,
			nil,
		},
		{

			"Case 4: Happy path, empty object for quota plan",
			`{}`,
			nil,
		},
		{
			"Case 5: Happy path, everything is valid with only filters and empty quotagroups",
			`{ "filters": [{ "attributeId": "13", "options": [ "18-99" ] }], "quotaGroups": [] }`,
			nil,
		},
		{
			"Case 6: Happy path, everything is valid with only quotagroups and empty filters",
			`{ "filters": [], "quotaGroups": [{ "name": "group 1", "quotaCells": [{ "quotaNodes": [{ "attributeId": "11", "options": [ "1" ] }], "count": 50 }, { "quotaNodes": [{ "attributeId": "11", "options": [ "2" ] }],"count": 50} ] }] }`,
			nil,
		},
		{
			"Case 7: Happy path, everything is valid with empty quotagroups and empty filters",
			`{ "filters": [], "quotaGroups": [] }`,
			nil,
		},
		{
			"Case 8: Error path, Empty quota cells",
			`{ "filters": [{ "attributeId": "13", "options": [ "18-99" ] }], "quotaGroups": [{ "name": "group 1", "quotaCells": [] }] }`,
			samplify.ErrMissingQuotaCells,
		},
		{
			"Case 9: Error path, nil quota cells",
			`{ "filters": [{ "attributeId": "13", "options": [ "18-99" ] }], "quotaGroups": [{ "name": "group 1" }] }`,
			samplify.ErrMissingQuotaCells,
		},
		{
			"Case 10: Error path, nil perc and count",
			`{ "filters": [{ "attributeId": "13", "options": [ "18-99" ] }], "quotaGroups": [{ "name": "group 1", "quotaCells": [{ "quotaNodes": [{ "attributeId": "11", "options": [ "1" ] }], "perc": 50 }, { "quotaNodes": [{ "attributeId": "11", "options": [ "2" ] }] } ] }] }`,
			samplify.ErrAllocationNotProvided,
		},
		{
			"Case 11: Error path, when both perc and count are provided",
			`{ "filters": [{ "attributeId": "13", "options": [ "18-99" ] }], "quotaGroups": [{ "name": "group 1", "quotaCells": [{ "quotaNodes": [{ "attributeId": "11", "options": [ "1" ] }], "perc": 50, "count": 50 }, { "quotaNodes": [{ "attributeId": "11", "options": [ "2" ] }],"perc": 50, "count": 50 } ] }] }`,
			samplify.ErrAmbigiuosAllocation,
		},
		{
			"Case 12: Error path, mismatched cells allocation type perc in one cell and count in other",
			`{ "filters": [{ "attributeId": "13", "options": [ "18-99" ] }], "quotaGroups": [{ "name": "group 1", "quotaCells": [{ "quotaNodes": [{ "attributeId": "11", "options": [ "1" ] }], "count": 50 }, { "quotaNodes": [{ "attributeId": "11", "options": [ "2" ] }],"perc": 50} ] }] }`,
			samplify.ErrInconsistentAllocationType,
		},
	}

	for _, table := range tables {
		qp := &samplify.QuotaPlan{}
		err := json.Unmarshal([]byte(table.input), qp)
		if err != nil {
			t.Fatal(err)
		}

		e := samplify.ValidateQuotaPlan(qp)
		if e != table.expected {
			t.Errorf("%s validation failed got: `%v`, want `%v` ", table.name, e, table.expected)
		}
	}
}
