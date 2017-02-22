package main

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type field struct {
	val       reflect.Value
	validator string
}

// Unpack populates the fields of the struct pointed to by ptr
// from the HTTP request parameters in req.
func Unpack(req *http.Request, ptr interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}

	// Build map of fields keyed by effective name.
	fields := make(map[string]field)
	v := reflect.ValueOf(ptr).Elem() // the struct variable
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) // a reflect.StructField
		tag := fieldInfo.Tag           // a reflect.StructTag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}

		fields[name] = field{v.Field(i), tag.Get("validate")}
	}

	// Update struct field for each parameter in the request.
	for name, values := range req.Form {
		f := fields[name].val
		if !f.IsValid() {
			continue // ignore unrecognized HTTP parameters
		}
		validator := fields[name].validator
		for _, value := range values {
			if f.Kind() == reflect.Slice {
				elem := reflect.New(f.Type().Elem()).Elem()
				if err := populate(elem, value, validator); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
				f.Set(reflect.Append(f, elem))
			} else {
				if err := populate(f, value, validator); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
			}
		}
	}
	return nil
}

func populate(v reflect.Value, value string, validator string) error {
	// validator specified
	if validator != "" {
		valid, err := validateByName(validator, value)
		if err != nil {
			return err
		}

		if !valid {
			return fmt.Errorf("invalid data format:%s", value)
		}
	}

	switch v.Kind() {
	case reflect.String:
		v.SetString(value)

	case reflect.Int:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		v.SetInt(i)

	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		v.SetBool(b)

	default:
		return fmt.Errorf("unsupported kind %s", v.Type())
	}
	return nil
}
