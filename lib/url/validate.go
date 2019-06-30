package url

import (
	"errors"
	"net/url"
	"strconv"
	"strings"
	"unicode/utf8"

	validator "github.com/asaskevich/govalidator"
)

const (
	//URLMaxLength is the max length of URL allowed by browsers
	URLMaxLength = 2083
	//URLMinLength is the min length of URL allowed by browsers
	URLMinLength = 3
)

var (
	// ErrURLBlank is error for blank URL
	ErrURLBlank = errors.New("the URL cannot be blank")
	// ErrURLMaxLength is error for MAX length (URLMaxLength) of URL
	ErrURLMaxLength = errors.New("the URL length cannot exceed " + strconv.Itoa(URLMaxLength) + " characters")
	// ErrURLMinLength is error for MAX length (URLMinLength) of URL
	ErrURLMinLength = errors.New("the URL length cannot be less than " + strconv.Itoa(URLMinLength) + " characters")
	// ErrURLPrefix is error for a prefix '.' in URL
	ErrURLPrefix = errors.New("the URL cannot have a prefix '.'")
	// ErrURLHostPrefix is error for a prefix '.' in URL host
	ErrURLHostPrefix = errors.New("the URL host cannot have a prefix '.'")
	// ErrURLHost is error for URL that has no host
	ErrURLHost = errors.New("the URL host is not present")
	// ErrURLFragment is error for URL with fragments
	ErrURLFragment = errors.New("the URL has a fragment(#) which is not supported")
	// ErrURLInvalid is error for invalid URL
	ErrURLInvalid = errors.New("the URL is invalid")
)

// IsURLValid checks if the string `str` is an URL
func IsURLValid(str string) (bool, error) {
	if str == "" {
		return false, ErrURLBlank
	}

	c := utf8.RuneCountInString(str)
	if c >= URLMaxLength {
		return false, ErrURLMaxLength
	}

	if c <= URLMinLength {
		return false, ErrURLMinLength
	}

	if strings.HasPrefix(str, ".") {
		return false, ErrURLPrefix
	}

	strTemp := str
	if strings.Contains(str, ":") && !strings.Contains(str, "://") {
		// support no indicated urlscheme but with colon for port number
		// http:// is appended so url.Parse will succeed, strTemp used so it does not impact validator.IsURL()
		strTemp = "http://" + str
	}
	u, err := url.Parse(strTemp)
	if err != nil {
		return false, err
	}

	if strings.HasPrefix(u.Host, ".") {
		return false, ErrURLHostPrefix
	}

	if u.Host == "" && (u.Path != "" && !strings.Contains(u.Path, ".")) {
		return false, ErrURLHost
	}

	if u.Fragment != "" {
		return false, ErrURLFragment
	}

	if !validator.IsURL(str) {
		return false, ErrURLInvalid
	}

	return true, nil
}
