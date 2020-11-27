package main

import (
	"errors"
	"reflect"
	"strings"
)

// OverloadOptions is ...
type OverloadOptions []OverloadCallback

// OverloadCallback is ...
type OverloadCallback struct {
	argTypes   []string
	returnType string
	callback   func(OverloadArgs) interface{}
}

// OverloadArgs is ...
type OverloadArgs []interface{}

// OverloadResponse is ...
type OverloadResponse struct {
	value        interface{}
	responseType string
}

// Overload is ...
func Overload(overloadList OverloadOptions, args ...interface{}) (results OverloadResponse, err error) {
	//* Panic if overload not specified
	if overloadList == nil || len(overloadList) == 0 {
		panic("Overload Error: No Overloads Specified")
	}

	//* Iterate over Overloads
	overloadFound := false
	for _, indexValue := range overloadList {

		//* Error If types are greater than args
		if len(args)+1 < len(indexValue.argTypes) {
			return OverloadResponse{}, errors.New("Overload Error: Length of types is greater than args")
		}

		//* Find Overload
		overloadFound = true
		for index, indexType := range indexValue.argTypes {
			if getType(args[index]) != strings.ToLower(indexType) {
				overloadFound = false
				break
			}
		}

		//* Run Callback once found
		if overloadFound {
			return OverloadResponse{
				value:        indexValue.callback(args),
				responseType: indexValue.returnType,
			}, nil
		}
	}

	//* Return Overload not found error
	return OverloadResponse{}, errors.New("Overload Error: Overload not found for args")
}

func getType(arg interface{}) (res string) {
	argType := reflect.TypeOf(arg)

	for argType.Kind() == reflect.Ptr {
		argType = argType.Elem()
		res += "*"
	}

	return res + argType.Name()

}

//* Main for linter purpose
func main() {}
