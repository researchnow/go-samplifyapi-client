package url

import (
	"net/url"
)

// Format returns the formatted TestLink appending the custom paramaters
func Format(u *url.URL, m map[string]string) (*string, error) {
	q := u.Query()

	for key, value := range m {
		q.Set(key, value)
	}
	u.RawQuery = q.Encode()

	decodedURL := u.String()
	res, err := url.QueryUnescape(decodedURL)
	return &res, err
}
