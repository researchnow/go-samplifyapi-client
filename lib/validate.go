package samplify

import (
	"errors"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/asaskevich/govalidator"
)

// Constants ...
const (
	URLMaxLength = 2083
	URLMinLength = 3
)

// Validation Errors
var (
	ErrRequiredFieldEmpty    = errors.New("required field is empty")
	ErrInvalidFieldValue     = errors.New("invalid field value")
	ErrInvalidQuotaCellValue = errors.New("invalid quota cell value, perc or count must be specified")

	// URL validation errros
	ErrURLBlank      = errors.New("the URL cannot be blank")
	ErrURLMaxLength  = errors.New("the URL length cannot exceed " + strconv.Itoa(URLMaxLength) + " characters")
	ErrURLMinLength  = errors.New("the URL length cannot be less than " + strconv.Itoa(URLMinLength) + " characters")
	ErrURLPrefix     = errors.New("the URL cannot have a prefix '.'")
	ErrURLHostPrefix = errors.New("the URL host cannot have a prefix '.'")
	ErrURLHost       = errors.New("the URL host is not present")
	ErrURLFragment   = errors.New("the URL has a fragment(#) which is not supported")
	ErrURLInvalid    = errors.New("the URL is invalid")
)

// Validate ...
func Validate(obj interface{}) error {
	kind := reflect.TypeOf(obj).Kind()
	switch kind {
	case reflect.Slice:
		o := reflect.ValueOf(obj)
		if o.Len() == 0 {
			return ErrRequiredFieldEmpty
		}
		for i := 0; i < o.Len(); i++ {
			if err := ValidateStruct(o.Index(i).Interface()); err != nil {
				return err
			}
		}
		return nil
	default:
		o := reflect.ValueOf(obj)
		return ValidateStruct(o.Interface())
	}
}

// ValidateStruct ...
func ValidateStruct(obj interface{}) error {
	valid, err := govalidator.ValidateStruct(obj)
	if err != nil {
		return err
	}
	if !valid {
		return ErrRequiredFieldEmpty
	}
	return nil
}

// ValidateNotNull ...
func ValidateNotNull(obj interface{}) error {
	if reflect.ValueOf(obj).IsNil() {
		return ErrRequiredFieldEmpty
	}
	return nil
}

// ValidateNotEmpty ...
func ValidateNotEmpty(values ...string) error {
	for _, val := range values {
		if len(val) == 0 {
			return ErrRequiredFieldEmpty
		}
	}
	return nil
}

// ValidateAction ...
func ValidateAction(action Action) error {
	if action == ActionClosed ||
		action == ActionLaunched ||
		action == ActionPaused {
		return nil
	}
	return ErrInvalidFieldValue
}

// ValidateEmail ...
func ValidateEmail(email ...string) error {
	for _, e := range email {
		if !govalidator.IsEmail(e) {
			return ErrInvalidFieldValue
		}
	}
	return nil
}

// ValidateLanguageISOCode ...
func ValidateLanguageISOCode(val string) error {
	if govalidator.IsISO693Alpha2(val) {
		return nil
	}
	return ErrInvalidFieldValue
}

// ValidateDeviceType ...
func ValidateDeviceType(val DeviceType) error {
	if val != DeviceTypeDesktop &&
		val != DeviceTypeMobile &&
		val != DeviceTypeTablet {
		return ErrInvalidFieldValue
	}
	return nil
}

// ValidateExclusionType ...
func ValidateExclusionType(val ExclusionType) error {
	if val != ExclusionTypeProject &&
		val != ExclusionTypeTag {
		return ErrInvalidFieldValue
	}
	return nil
}

// ValidateDeliveryType ...
func ValidateDeliveryType(val DeliveryType) error {
	if val != DeliveryTypeSlow &&
		val != DeliveryTypeBalanced &&
		val != DeliveryTypeFast {
		return ErrInvalidFieldValue
	}
	return nil
}

// ValidateQuotaPlan ...
func ValidateQuotaPlan(val *QuotaPlan) error {
	if val != nil &&
		val.Filters == nil &&
		val.QuotaGroups == nil {
		return ErrInvalidFieldValue
	}

	//quota cell must have either percentage or count
	for _, vq := range val.QuotaGroups {
		if vq.QuotaCells == nil {
			return ErrInvalidFieldValue
		}
		for _, vc := range vq.QuotaCells {
			if vc.Perc == nil && vc.Count == nil {
				return ErrInvalidQuotaCellValue
			}
		}
	}
	return nil
}

// ValidateSurveyURL ...
func ValidateSurveyURL(baseURL string) error {
	if baseURL == "" {
		return ErrURLBlank
	}

	c := utf8.RuneCountInString(baseURL)
	if c >= URLMaxLength {
		return ErrURLMaxLength
	}

	if c <= URLMinLength {
		return ErrURLMinLength
	}

	if strings.HasPrefix(baseURL, ".") {
		return ErrURLPrefix
	}

	strTemp := baseURL
	if strings.Contains(baseURL, ":") && !strings.Contains(baseURL, "://") {
		// support no indicated urlscheme but with colon for port number
		// http:// is appended so url.Parse will succeed, strTemp used so it does not impact validator.IsURL()
		strTemp = "http://" + baseURL
	}
	u, err := url.Parse(strTemp)
	if err != nil {
		return err
	}

	if strings.HasPrefix(u.Host, ".") {
		return ErrURLHostPrefix
	}

	if u.Host == "" && (u.Path != "" && !strings.Contains(u.Path, ".")) {
		return ErrURLHost
	}

	if u.Fragment != "" || strings.Contains(baseURL, "#") {
		return ErrURLFragment
	}

	if !govalidator.IsURL(baseURL) {
		return ErrURLInvalid
	}
	return nil
}

// IsCountryCodeOrEmpty ...
func IsCountryCodeOrEmpty(countryCode string) error {
	if len(countryCode) > 0 && !govalidator.IsISO3166Alpha2(countryCode) {
		return ErrInvalidFieldValue
	}
	return nil
}

// IsLanguageCodeOrEmpty ...
func IsLanguageCodeOrEmpty(languageCode string) error {
	if len(languageCode) > 0 && !govalidator.IsISO693Alpha2(languageCode) {
		return ErrInvalidFieldValue
	}
	return nil
}

func init() {
	govalidator.CustomTypeTagMap.Set("languageISOCode", govalidator.CustomTypeValidator(func(i interface{}, o interface{}) bool {
		var err error
		switch v := i.(type) {
		case string:
			err = ValidateLanguageISOCode(v)
		case *string:
			err = ValidateLanguageISOCode(*v)
		case []string:
			for _, code := range v {
				if ValidateLanguageISOCode(code) != nil {
					return false
				}
			}
			return true
		case *[]string:
			for _, code := range *v {
				if ValidateLanguageISOCode(code) != nil {
					return false
				}
			}
			return true
		default:
			return false
		}
		if err != nil {
			return false
		}
		return true
	}))
	govalidator.CustomTypeTagMap.Set("DeviceType", govalidator.CustomTypeValidator(func(i interface{}, o interface{}) bool {
		var err error
		switch v := i.(type) {
		case DeviceType:
			err = ValidateDeviceType(v)
		case *DeviceType:
			err = ValidateDeviceType(*v)
		case []DeviceType:
			for _, device := range v {
				if ValidateDeviceType(device) != nil {
					return false
				}
			}
			return true
		case *[]DeviceType:
			for _, device := range *v {
				if ValidateDeviceType(device) != nil {
					return false
				}
			}
			return true
		default:
			return false
		}
		if err != nil {
			return false
		}
		return true
	}))
	govalidator.CustomTypeTagMap.Set("ExclusionType", govalidator.CustomTypeValidator(func(i interface{}, o interface{}) bool {
		var err error
		switch v := i.(type) {
		case ExclusionType:
			err = ValidateExclusionType(v)
		case []ExclusionType:
			for _, ex := range v {
				if ValidateExclusionType(ex) != nil {
					return false
				}
			}
			return true
		default:
			return false
		}
		if err != nil {
			return false
		}
		return true
	}))
	govalidator.CustomTypeTagMap.Set("DeliveryType", govalidator.CustomTypeValidator(func(i interface{}, o interface{}) bool {
		var err error
		switch v := i.(type) {
		case DeliveryType:
			err = ValidateDeliveryType(v)
		case *DeliveryType:
			err = ValidateDeliveryType(*v)
		default:
			return false
		}
		if err != nil {
			return false
		}
		return true
	}))
	govalidator.CustomTypeTagMap.Set("quotaPlan", govalidator.CustomTypeValidator(func(i interface{}, o interface{}) bool {
		var err error
		switch v := i.(type) {
		case QuotaPlan:
			err = ValidateQuotaPlan(&v)
		case *QuotaPlan:
			err = ValidateQuotaPlan(v)
		default:
			return false
		}
		if err != nil {
			return false
		}
		return true
	}))
	govalidator.CustomTypeTagMap.Set("surveyURL", govalidator.CustomTypeValidator(func(i interface{}, o interface{}) bool {
		var err error
		switch v := i.(type) {
		case string:
			err = ValidateSurveyURL(v)
		case *string:
			err = ValidateSurveyURL(*v)
		default:
			return false
		}
		if err != nil {
			return false
		}
		return true
	}))
}
