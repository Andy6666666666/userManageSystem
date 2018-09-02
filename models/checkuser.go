package models

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

const tagName = "validata"

var mobileRe = regexp.MustCompile(`^1([38][0-9]|14[57]|5[^4])\d{8}$`)
var nameOrPwdRe = regexp.MustCompile(`^[a-zA-Z0-9_-]{4,16}$`)

type Validator interface {
	Validate(interface{}) (bool, error)
}

type DefaultValidator struct {
}

func (v DefaultValidator) Validate(val interface{}) (bool, error) {
	return true, nil
}

type StringValidator struct {
	Min int
	Max int
}

func (v StringValidator) Validate(val interface{}) (bool, error) {

	l := len(val.(string))

	if l == 0 {
		return false, fmt.Errorf("cannot be blank")
	}

	if l < v.Min {
		return false, fmt.Errorf("should be at least %v chars long", v.Min)

	}

	if v.Max >= v.Min && l > v.Max {
		return false, fmt.Errorf("should be less than %v chars long", v.Max)
	}

	return true, nil
}

type MobileValidator struct {
}

func (v MobileValidator) Validate(val interface{}) (bool, error) {

	if !mobileRe.MatchString(val.(string)) {
		return false, fmt.Errorf("is not a valid mobile number")
	}

	return true, nil
}

type NameOrPwdValidator struct {
}

func (v NameOrPwdValidator) Validate(val interface{}) (bool, error) {

	if !nameOrPwdRe.MatchString(val.(string)) {
		return false, fmt.Errorf("is not a valid name or password")
	}

	return true, nil
}

func getValidatorFromTag(tag string) Validator {

	args := strings.Split(tag, ",")

	fmt.Printf("args=%v", args)

	switch args[0] {

	case "string":

		validator := StringValidator{}

		fmt.Sscanf(strings.Join(args[1:], ","), "min=%d,max=%d", &validator.Min, &validator.Max)

		return validator

	case "nameOrPwd":
		return NameOrPwdValidator{}

	case "mobile":
		return MobileValidator{}
	}

	return DefaultValidator{}

}

func validateStruct(s interface{}) []error {

	errs := []error{}

	v := reflect.ValueOf(s)

	for i := 0; i < v.NumField(); i++ {

		//利用反射获取structTag

		tag := v.Type().Field(i).Tag.Get(tagName)

		if tag == "" {
			continue
		}

		validator := getValidatorFromTag(tag)

		valid, err := validator.Validate(v.Field(i).Interface())

		if !valid && err != nil {
			fmt.Printf(err.Error())
			errs = append(errs, fmt.Errorf("%s %s", v.Type().Field(i).Name, err.Error()))

		}

	}

	return errs
}
