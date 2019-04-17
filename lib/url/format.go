package url

import (
	"fmt"
	"net/url"
	"strings"
)

const (
	//TemplatePSID Default Paramter value of PSID
	TemplatePSID = "<#IdParameter[Value]>"
	//TemplatePID Default Paramter value of PID
	TemplatePID = "<#DubKnowledge[1500/Entity id]>"
	//TemplateSecurityKey Default Paramter value of K2
	TemplateSecurityKey = "<#Project[Secure Key 2]>"
)

// CustomParams is a type of map that will hold the parameter value as key and parameter template as value.
type CustomParams map[string]string

// Format returns the formatted TestLink appending the custom paramaters
func Format(u *url.URL, m CustomParams) string {
	queryString := ""
	queryValues := u.Query()
	for key := range queryValues {
		queryString = fmt.Sprintf("%s&%s=%s", queryString, key, queryValues.Get(key))
	}
	for key, value := range m {
		queryString = fmt.Sprintf("%s&%s=%s", queryString, key, value)
	}
	query := strings.Trim(queryString, queryString[:1])
	u.RawQuery = query
	return u.String()
}
