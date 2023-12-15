package typeconv

import (
	"github.com/mpyw/go-snippets/reflectx"
	"golang.org/x/exp/constraints"
)

// CoerceIntMap interconverts map[T]int types of arbitrary precision
func CoerceIntMap[TReturn constraints.Integer, TArg constraints.Integer, K comparable](m map[K]TArg) map[K]TReturn {
	newMap := make(map[K]TReturn, len(m))
	for key, value := range m {
		newMap[key] = TReturn(value)
	}
	return newMap
}

// AsAnyMap converts map[K]T into map[K]any
func AsAnyMap[T any, K comparable](m map[K]T) map[K]any {
	newMap := make(map[K]any, len(m))
	for key, value := range m {
		newMap[key] = value
	}
	return newMap
}

// AsAnyMapCoercingNils converts map[K]T into map[K]any, coercing typed nils into non-typed nils
func AsAnyMapCoercingNils[T any, K comparable](m map[K]T) map[K]any {
	newMap := make(map[K]any, len(m))
	for key, value := range m {
		if reflectx.IsNil(value) {
			newMap[key] = nil
		} else {
			newMap[key] = value
		}
	}
	return newMap
}
