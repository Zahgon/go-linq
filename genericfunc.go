package linq

import (
	"reflect"
)

// genericType represents a any reflect.Type.
type genericType int

var genericTp = reflect.TypeOf(new(genericType)).Elem()

// functionCache keeps genericFunc reflection objects in cache.
type functionCache struct {
	MethodName string
	ParamName  string
	FnValue    reflect.Value
	FnType     reflect.Type
	TypesIn    []reflect.Type
	TypesOut   []reflect.Type
}

// genericFunc is a type used to validate and call dynamic functions.
type genericFunc struct {
	Cache *functionCache
}

// Call calls a dynamic function.
func (g *genericFunc) Call(params ...any) any { _ = "STUB: not implemented"; return *new(any) }

// newGenericFunc instantiates a new genericFunc pointer
func newGenericFunc(methodName, paramName string, fn any, validateFunc func(*functionCache) error) (*genericFunc, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// simpleParamValidator creates a function to validate genericFunc based in the
// In and Out function parameters.
func simpleParamValidator(In []reflect.Type, Out []reflect.Type) func(cache *functionCache) error {
	_ = "STUB: not implemented"
	return nil
}

// newElemTypeSlice creates a slice of items elem types.
func newElemTypeSlice(items ...any) []reflect.Type { _ = "STUB: not implemented"; return nil }

// formatFnSignature formats the func signature based in the parameters types.
func formatFnSignature(In []reflect.Type, Out []reflect.Type) string {
	_ = "STUB: not implemented"
	return ""
}
