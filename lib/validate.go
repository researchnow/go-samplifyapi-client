package samplify

import (
	"errors"
	"reflect"
	"strings"

	"github.com/asaskevich/govalidator"
)

// Validation Errors
var (
	ErrRequiredFieldEmpty = errors.New("required field is empty")
	ErrInvalidFieldValue  = errors.New("invalid field value")
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
	return nil
}

// ValidateSurveyURL ...
func ValidateSurveyURL(val string) error {
	s := strings.Split(val, "?")
	if len(s) != 2 {
		return ErrInvalidFieldValue
	}
	yes := govalidator.IsURL(s[0])
	if yes {
		return nil
	}
	return ErrInvalidFieldValue
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
