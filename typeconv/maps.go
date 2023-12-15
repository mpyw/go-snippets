package typeconv

import (
	"github.com/mpyw/go-snippets/reflectx"
	"golang.org/x/exp/constraints"
)

// CoerceIntSlice interconverts []int types of arbitrary precision
func CoerceIntSlice[TReturn constraints.Integer, TArg constraints.Integer](slice []TArg) []TReturn {
	var newSlice []TReturn
	for _, elem := range slice {
		newSlice = append(newSlice, TReturn(elem))
	}
	return newSlice
}

// AsAnySlice converts []T into []any
func AsAnySlice[T any](slice []T) []any {
	var newSlice []any
	for _, elem := range slice {
		newSlice = append(newSlice, elem)
	}
	return newSlice
}

// AsAnySliceCoercingNils converts []T into []any, coercing typed nils into non-typed nils
func AsAnySliceCoercingNils[T any](slice []T) []any {
	var newSlice []any
	for _, elem := range slice {
		if reflectx.IsNil(elem) {
			newSlice = append(newSlice, nil)
		} else {
			newSlice = append(newSlice, elem)
		}
	}
	return newSlice
}
