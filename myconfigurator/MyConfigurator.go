package myconfigurator

import (
	"errors"
	"reflect"
)

const (
	CUSTOM uint8 = iota
	JSON
	XML
)

var wrongTypeError = errors.New("type must be a pointer to a struct")

func GetConfiguration(confType uint8, obj interface{}, filename string) (err error) {
	switch confType {
	case CUSTOM:
		mysRValue := reflect.ValueOf(obj)
		// Check if type is pointer.
		if mysRValue.Kind() != reflect.Ptr || mysRValue.IsNil() {
			return wrongTypeError
		}
		// Get and confirm the struct value.
		mysRValue = mysRValue.Elem()
		// *object => object
		// reflection value of *object .Elem() => object() (Settable)
		if mysRValue.Kind() != reflect.Struct {
			return wrongTypeError
		}
		err = MarshalCustomConfig(mysRValue, filename)
	case JSON:
		err = decodeJSONConfig(obj, filename)
	case XML:
		err = decodeXMLConfig(obj, filename)
	}

	return err
}
