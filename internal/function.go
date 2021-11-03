package internal

import (
	"github.com/celicoo/ttest/internal/internal"
	"path"
	"reflect"
	"runtime"
)

// NewFunction returns the address of a Function instance. (*Function)(nil) is
// returned if i is not a function.
func NewFunction(i interface{}) *Function {
	v := reflect.ValueOf(i)
	if v.Kind() != reflect.Func {
		return nil
	}
	return &Function{v}
}

// Function represents a Go function.
type Function struct {
	reflect.Value
}

// Call calls the function f with the input arguments and returns the returning
// values as a slice of interface{} values.
// Call panics if the call fails.
func (f Function) Call(arguments []interface{}) []interface{} {
	var vv []reflect.Value
	for i := range arguments {
		i := arguments[i]
		vv = append(
			vv,
			reflect.ValueOf(i),
		)
	}
	return internal.InterfacesOf(f.Value.Call(vv)...)
}

// Name returns the basename of f.
func (f Function) Name() string {
	p := f.Pointer()
	s := runtime.FuncForPC(p).Name()
	return path.Base(s)
}
