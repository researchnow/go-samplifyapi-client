package url

import (
	"fmt"
	"net/url"
	"sort"
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
	u.ForceQuery = true
	urlstring := URLString(u)

	if urlstring == "" {
		return urlstring
	}

	keys := sortMapKeys(m)
	// there are some url params in the url itself.

	index := strings.Index(urlstring, "?")

	if urlstring[len(urlstring)-1:] != "?" && index < 0 {
		urlstring = fmt.Sprintf("%s?", urlstring)
	}

	if index > -1 {
		urlstring = fmt.Sprintf("%s&", urlstring)
	}

	for _, key := range keys {
		urlstring = fmt.Sprintf("%s%s=%s&", urlstring, key, m[key])
	}

	urlstring = strings.Trim(urlstring, "&")
	return urlstring
}

func sortMapKeys(m CustomParams) []string {
	keys := make([]string, 0, len(m))

	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	return keys
}

// **** Copied this from the default url package to ensure the messed up
// <#IdParameter[Value]> stuff works. so, the encode on the host and the fragment are removed.

// URLString reassembles the URL into a valid URL string.
// The general form of the result is one of:
//
//	scheme:opaque?query#fragment
//	scheme://userinfo@host/path?query#fragment
//
// If u.Opaque is non-empty, String uses the first form;
// otherwise it uses the second form.
// Any non-ASCII characters in host are escaped.
// To obtain the path, String uses u.EscapedPath().
//
// In the second form, the following rules apply:
//	- if u.Scheme is empty, scheme: is omitted.
//	- if u.User is nil, userinfo@ is omitted.
//	- if u.Host is empty, host/ is omitted.
//	- if u.Scheme and u.Host are empty and u.User is nil,
//	   the entire scheme://userinfo@host/ is omitted.
//	- if u.Host is non-empty and u.Path begins with a /,
//	   the form host/path does not add its own /.
//	- if u.RawQuery is empty, ?query is omitted.
//	- if u.Fragment is empty, #fragment is omitted.
func URLString(u *url.URL) string {
	var buf strings.Builder
	if u.Scheme != "" {
		buf.WriteString(u.Scheme)
		buf.WriteByte(':')
	}
	if u.Opaque != "" {
		buf.WriteString(u.Opaque)
	} else {
		if u.Scheme != "" || u.Host != "" || u.User != nil {
			if u.Host != "" || u.Path != "" || u.User != nil {
				buf.WriteString("//")
			}
			if ui := u.User; ui != nil {
				buf.WriteString(ui.String())
				buf.WriteByte('@')
			}
			if h := u.Host; h != "" {
				buf.WriteString(h)
			}
		}
		path := u.EscapedPath()
		if path != "" && path[0] != '/' && u.Host != "" {
			buf.WriteByte('/')
		}
		if buf.Len() == 0 {
			// RFC 3986 §4.2
			// A path segment that contains a colon character (e.g., "this:that")
			// cannot be used as the first segment of a relative-path reference, as
			// it would be mistaken for a scheme name. Such a segment must be
			// preceded by a dot-segment (e.g., "./this:that") to make a relative-
			// path reference.
			if i := strings.IndexByte(path, ':'); i > -1 && strings.IndexByte(path[:i], '/') == -1 {
				buf.WriteString("./")
			}
		}
		buf.WriteString(path)
	}
	if u.RawQuery != "" {
		buf.WriteByte('?')
		buf.WriteString(u.RawQuery)
	}
	if u.Fragment != "" {
		buf.WriteByte('#')
		buf.WriteString(u.Fragment)
	}
	return buf.String()
}
