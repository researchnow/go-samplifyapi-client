package url_test

import (
	"net/url"
	"sort"
	"testing"

	formatURL "github.com/researchnow/go-samplifyapi-client/lib/url"
)

func TestFormat(t *testing.T) {
	tables := []struct {
		name        string
		inputURL    string
		inputParams formatURL.CustomParams
		output      string
	}{
		{
			"Case 1: Happy case - Passing all the parameters ",
			"http://www.google.com",
			formatURL.CustomParams{"psid": "{{PSID}}", "pid": "{{PID}}", "k2": "{{K2}}"},
			"http://www.google.com?k2={{K2}}&pid={{PID}}&psid={{PSID}}",
		},
		{
			"Case 2: Happy case - Passing all the parameters ",
			"www.google.com",
			formatURL.CustomParams{"psid": "{{PSID}}", "pid": "{{PID}}", "k2": "{{K2}}"},
			"www.google.com?k2={{K2}}&pid={{PID}}&psid={{PSID}}",
		},
		{
			"Case 3: Passing Few Parameters ",
			"http://www.google.com",
			formatURL.CustomParams{"k2": "{{K2}}", "trackerID": "{{PSID}}", "projectID": "{{PID}}"},
			"http://www.google.com?k2={{K2}}&projectID={{PID}}&trackerID={{PSID}}",
		},
		{
			"Case 4: Passing Few Parameters ",
			"http://www.google.com",
			formatURL.CustomParams{"security": "{{K2}}", "tracker": "{{PSID}}", "project": "{{PID}}"},
			"http://www.google.com?project={{PID}}&security={{K2}}&tracker={{PSID}}",
		},
	}
	for _, table := range tables {
		url, _ := url.Parse(table.inputURL)
		actual := formatURL.Format(url, sortMapKeys(table.inputParams))
		if actual != table.output {
			t.Errorf("%s validation failed got: %v, want %v ", table.name, actual, table.output)
		}
	}

}

func sortMapKeys(m formatURL.CustomParams) formatURL.CustomParams {
	keys := make([]string, 0, len(m))
	result := formatURL.CustomParams{}
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		result[key] = m[key]
	}
	return result
}
