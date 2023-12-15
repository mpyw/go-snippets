// Package mocking provides utilities for https://github.com/uber-go/mock
package mocking

import (
	"fmt"

	"go.uber.org/mock/gomock"
)

// FuncMatcher provides an instant functional matcher.
func FuncMatcher[T any, R bool | string](fn func(actual T) R) gomock.Matcher {
	return &funcMatcher[T, R]{fn: fn}
}

var _ gomock.Matcher = (*funcMatcher)(nil)

type funcMatcher[T any, R bool | string] struct {
	fn      func(actual T) R
	message string
}

func (m *funcMatcher[T, R]) Matches(actual any) bool {
	t, ok := actual.(T)
	if !ok {
		var typ T
		m.message = fmt.Sprintf("must be assignable to %T", typ)
		return false
	}
	switch r := any(m.fn(t)).(type) {
	case bool:
		if r {
			return true
		}
		m.message = "doesn't satisfy required constraint"
		return false
	case string:
		if r == "" {
			return true
		}
		m.message = r
		return false
	default:
		return false
	}
}

func (m *funcMatcher[T, R]) String() string {
	return m.message
}
