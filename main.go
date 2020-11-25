package main

import (
	"errors"
	"reflect"
)

// OverloadList is ...
type OverloadList []OverloadIndex

// OverloadIndex is ...
type OverloadIndex struct {
	types    []interface{}
	callback func(OverloadArgs) interface{}
}

// OverloadArgs is ...
type OverloadArgs []interface{}

// Overload is ...
func Overload(overloadList OverloadList, args ...interface{}) (results interface{}, err error) {
	//* Panic if overload not specified
	if overloadList == nil || len(overloadList) == 0 {
		panic("Overload Error: No Overloads Specified")
	}

	//* Iterate over Overloads
	overloadFound := false
	for _, indexValue := range overloadList {

		//* Error If types are greater than args
		if len(args)+1 < len(indexValue.types) {
			return nil, errors.New("Overload Error: Length of types is greater than args")
		}

		//* Find Overload
		overloadFound = true
		for index, indexType := range indexValue.types {
			if reflect.TypeOf(args[index]) != reflect.TypeOf(indexType) {
				overloadFound = false
				break
			}
		}

		//* Run Callback once found
		if overloadFound {
			return indexValue.callback(args), nil
		}
	}

	//* Return Overload not found error
	return nil, errors.New("Overload Error: Overload not found for args")
}

//* Main for linter purpose
func main() {}
