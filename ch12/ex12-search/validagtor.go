package main

import (
	"fmt"
	"reflect"
	"regexp"
)

var (
	// poor validation
	matcherEmail      = regexp.MustCompile("^[^@]+@[^@]+$")
	matcherCreditCard = regexp.MustCompile("^[0-9]{16}$")
)

type validator struct{}

func (v *validator) CheckEmail(param interface{}) bool {
	return matcherEmail.MatchString(fmt.Sprint(param))
}

func (v *validator) CheckCreditCard(param interface{}) bool {
	return matcherCreditCard.MatchString(fmt.Sprint(param))
}

func validateByName(name string, target interface{}) (bool, error) {
	v := reflect.ValueOf(target)
	var p validator
	f := reflect.ValueOf(&p).MethodByName(name)
	if !f.IsValid() {
		return false, fmt.Errorf("invalid name:%s", name)
	}
	values := f.Call([]reflect.Value{v})
	if len(values) == 1 && values[0].Kind() == reflect.Bool {
		return values[0].Bool(), nil
	}

	return false, fmt.Errorf("invalid name:" + name)
}
