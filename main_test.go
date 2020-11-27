package main

import (
	"fmt"
	"testing"
)

func TestOverload(t *testing.T) {
	var tests = []struct {
		a, b interface{}
		want interface{}
		err  error
	}{
		{0, 1, 1, nil},
		{1, 0, 1, nil},
		{2, -2, 0, nil},
		{0, -1, -1, nil},
		{-1, 0, -1, nil},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf(" %d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans, err := wrapperFunc(tt.a, tt.b)
			if tt.err != err {
				t.Errorf("got error %e when expecting %e", err, tt.err)
			} else if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}

	//* Int test
	wrapperFunc(1, 2)
}

//* Test Wrappers
func wrapperFunc(args ...interface{}) (int, error) {
	defs := OverloadOptions{
		OverloadCallback{
			argTypes:   []string{"int", "int"},
			returnType: "int",
			callback:   intCallback,
		},
	}

	res, err := Overload(defs, args)
	return ToInt(res.value), err
}

func panicWrapper() (interface{}, error) {
	return Overload(nil, nil)
}

func emptyDefWrapper() (interface{}, error) {
	defs := OverloadOptions{}
	return Overload(nil, defs)
}

//* Callbacks
func intCallback(args OverloadArgs) interface{} {
	return ToInt(args[0]) + ToInt(args[1])
}
