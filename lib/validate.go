package samplify

import (
	"errors"
	"reflect"

	"github.com/asaskevich/govalidator"
)

var errRequiredFieldEmpty = errors.New("required field is empty")
var errInvalidFieldValue = errors.New("invalid field value")

func validateStruct(obj interface{}) error {
	valid, err := govalidator.ValidateStruct(obj)
	if err != nil {
		return err
	}
	if !valid {
		return errRequiredFieldEmpty
	}
	return nil
}

func validateNotNull(obj interface{}) error {
	if reflect.ValueOf(obj).IsNil() {
		return errRequiredFieldEmpty
	}
	return nil
}

func validateNotEmpty(values ...string) error {
	for _, val := range values {
		if len(val) == 0 {
			return errRequiredFieldEmpty
		}
	}
	return nil
}

func validate(obj interface{}) error {
	kind := reflect.TypeOf(obj).Kind()
	switch kind {
	case reflect.Slice:
		o := reflect.ValueOf(obj)
		if o.Len() == 0 {
			return errRequiredFieldEmpty
		}
		for i := 0; i < o.Len(); i++ {
			if err := validateStruct(o.Index(i).Interface()); err != nil {
				return err
			}
		}
		return nil
	default:
		o := reflect.ValueOf(obj)
		return validateStruct(o.Interface())
	}
}

func validateAction(action Action) error {
	if action == ActionClosed ||
		action == ActionLaunched ||
		action == ActionPaused {
		return nil
	}
	return errInvalidFieldValue
}

func validateEmail(email ...string) error {
	for _, e := range email {
		if !govalidator.IsEmail(e) {
			return errInvalidFieldValue
		}
	}
	return nil
}

func validateDeviceType(device ...DeviceType) error {
	for _, d := range device {
		if d != DeviceTypeDesktop &&
			d != DeviceTypeMobile &&
			d != DeviceTypeTablet {
			return errInvalidFieldValue
		}
	}
	return nil
}

func isCountryCodeOrEmpty(countryCode string) error {
	if len(countryCode) > 0 && !govalidator.IsISO3166Alpha2(countryCode) {
		return errInvalidFieldValue
	}
	return nil
}

func isLanguageCodeOrEmpty(languageCode string) error {
	if len(languageCode) > 0 && !govalidator.IsISO693Alpha2(languageCode) {
		return errInvalidFieldValue
	}
	return nil
}

func isURLOrEmpty(url ...string) error {
	for _, u := range url {
		if len(u) > 0 && !govalidator.IsURL(u) {
			return errInvalidFieldValue
		}
	}
	return nil
}

func init() {
	govalidator.CustomTypeTagMap.Set("languageISOCode", govalidator.CustomTypeValidator(func(i interface{}, o interface{}) bool {
		switch v := i.(type) {
		case string:
			return govalidator.IsISO693Alpha2(v)
		case *string:
			return govalidator.IsISO693Alpha2(*v)
		case []string:
			for _, code := range v {
				if !govalidator.IsISO693Alpha2(code) {
					return false
				}
			}
			return true
		case *[]string:
			for _, code := range *v {
				if !govalidator.IsISO693Alpha2(code) {
					return false
				}
			}
			return true
		default:
			return false
		}
	}))
	govalidator.CustomTypeTagMap.Set("DeviceType", govalidator.CustomTypeValidator(func(i interface{}, o interface{}) bool {
		var isValid = func(val DeviceType) bool {
			if val != DeviceTypeDesktop &&
				val != DeviceTypeMobile &&
				val != DeviceTypeTablet {
				return false
			}
			return true
		}
		switch v := i.(type) {
		case DeviceType:
			return isValid(v)
		case *DeviceType:
			return isValid(*v)
		case []DeviceType:
			for _, device := range v {
				if !isValid(device) {
					return false
				}
			}
			return true
		case *[]DeviceType:
			for _, device := range *v {
				if !isValid(device) {
					return false
				}
			}
			return true
		default:
			return false
		}
	}))
	govalidator.CustomTypeTagMap.Set("ExclusionType", govalidator.CustomTypeValidator(func(i interface{}, o interface{}) bool {
		var isValid = func(val ExclusionType) bool {
			if val != ExclusionTypeProject &&
				val != ExclusionTypeTag {
				return false
			}
			return true
		}
		switch v := i.(type) {
		case ExclusionType:
			return isValid(v)
		case []ExclusionType:
			for _, ex := range v {
				if !isValid(ex) {
					return false
				}
			}
			return true
		default:
			return false
		}
	}))
	govalidator.CustomTypeTagMap.Set("DeliveryType", govalidator.CustomTypeValidator(func(i interface{}, o interface{}) bool {
		var isValid = func(val DeliveryType) bool {
			if val != DeliveryTypeSlow &&
				val != DeliveryTypeBalanced &&
				val != DeliveryTypeFast {
				return false
			}
			return true
		}
		switch v := i.(type) {
		case DeliveryType:
			return isValid(v)
		case *DeliveryType:
			return isValid(*v)
		default:
			return false
		}
	}))
}
