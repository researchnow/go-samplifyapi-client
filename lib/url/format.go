package url

import (
	"net/url"
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
func Format(u *url.URL, m CustomParams) (*string, error) {
	q := u.Query()

	for key, value := range m {
		q.Set(key, value)
	}
	u.RawQuery = q.Encode()

	decodedURL := u.String()
	res, err := url.QueryUnescape(decodedURL)
	return &res, err
}
