package url_test

import (
	"net/url"
	"testing"

	formatURL "github.com/morningconsult/go-samplifyapi-client/lib/url"
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
		{
			"Case 5: No custom parameters",
			"http://www.google.com",
			nil,
			"http://www.google.com?",
		},
		{
			"Case 6: url with parameters",
			"http://www.google.com?myparam=blah",
			nil,
			"http://www.google.com?myparam=blah",
		},
		{
			"Case 7: url with DK parameter",
			"http://www.google.com/testpath?myparam=<#IdParameter[Value]>",
			nil,
			"http://www.google.com/testpath?myparam=<#IdParameter[Value]>",
		},
		{
			"Case 8: url with DK multiple parameters  ",
			"http://www.google.com/testpath?myparam=<#IdParameter[Value]>&myparam2=<#DubKnowledge[1500/Entity id]>",
			nil,
			"http://www.google.com/testpath?myparam=<#IdParameter[Value]>&myparam2=<#DubKnowledge[1500/Entity id]>",
		},
		{
			"Case 9: url with DK multiple parameters & custom paramters ",
			"http://www.google.com/testpath?myparam=<#IdParameter[Value]>&myparam2=<#DubKnowledge[1500/Entity id]>",
			formatURL.CustomParams{"k2": "<#Project[Secure Key 2]>", "trackerID": "<#IdParameter[Value]>", "projectID": "<#DubKnowledge[1500/Entity id]>"},
			"http://www.google.com/testpath?myparam=<#IdParameter[Value]>&myparam2=<#DubKnowledge[1500/Entity id]>&k2=<#Project[Secure Key 2]>&projectID=<#DubKnowledge[1500/Entity id]>&trackerID=<#IdParameter[Value]>",
		},
		{
			"Case 9: url with DK multiple parameters & same custom paramters  ",
			"http://www.google.com/testpath?psid=<#IdParameter[Value]>&pid=<#DubKnowledge[1500/Entity id]>",
			formatURL.CustomParams{"k2": "<#Project[Secure Key 2]>", "psid": "<#IdParameter[Value]>", "pid": "<#DubKnowledge[1500/Entity id]>"},
			"http://www.google.com/testpath?psid=<#IdParameter[Value]>&pid=<#DubKnowledge[1500/Entity id]>&k2=<#Project[Secure Key 2]>&pid=<#DubKnowledge[1500/Entity id]>&psid=<#IdParameter[Value]>",
		},
	}
	for _, table := range tables {
		url, _ := url.Parse(table.inputURL)
		actual := formatURL.Format(url, table.inputParams)
		if actual != table.output {
			t.Errorf("%s validation failed got: %v, want %v ", table.name, actual, table.output)
		}
	}

}
