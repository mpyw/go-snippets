package typeconv

import (
	"fmt"
	"strconv"

	"golang.org/x/exp/constraints"
)

// Ref returns a pointer to the passed value
func Ref[T any](v T) *T {
	return &v
}

// IndirectOr unwraps a pointer value, falling back to default.
func IndirectOr[T any](v *T, defaultValue T) T {
	if v == nil {
		return defaultValue
	}
	return *v
}

// BoolTo converts bool to int
func BoolTo[T constraints.Integer](v bool) T {
	if v {
		return 1
	} else {
		return 0
	}
}

// AsBool converts T to bool
func AsBool[T comparable](v T) bool {
	var empty T
	return v != empty
}

// CoerceIntRef interconverts *int types of arbitrary precision
func CoerceIntRef[TReturn constraints.Integer, TArg constraints.Integer](v *TArg) *TReturn {
	if v == nil {
		return nil
	}
	return Ref(TReturn(*v))
}

// ItoaRef converts *int into *string
func ItoaRef(v *int) *string {
	if v == nil {
		return nil
	}
	return Ref(strconv.Itoa(*v))
}

// AtoiRef converts *string into *int
func AtoiRef(v *string) (*int, error) {
	if v == nil {
		return nil, nil
	}
	i, err := strconv.Atoi(*v)
	if err != nil {
		return nil, fmt.Errorf("typeconv.AtoiRef: %w", err)
	}
	return &i, nil
}
