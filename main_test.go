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
func wrapperFunc(args ...interface{}) (interface{}, error) {
	defs := OverloadList{
		OverloadIndex{
			[]interface{}{1, 1},
			intCallback,
		},
	}
	return Overload(defs, args)
}

func panicWrapper() (interface{}, error) {
	return Overload(nil, nil)
}

func emptyDefWrapper() (interface{}, error) {
	defs := OverloadList{}
	return Overload(nil, defs)
}

//* Callbacks
func intCallback(args OverloadArgs) interface{} {
	return args[0].(int) + args[1].(int)
}
