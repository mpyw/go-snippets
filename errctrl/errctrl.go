// Package errctrl is intended to simplify error checking and assignments to temporary variables.
package errctrl

// MustExec panics if the first arg is an error.
//   - Allowed in production code only if it never fails.
//   - Don't use in production code if it is likely to fail. Use only in tests.
func MustExec(err error) {
	if err != nil {
		panic(err)
	}
}

// Must panics if the second arg is an error.
//   - Allowed in production code only if it never fails.
//   - Don't use in production code if it is likely to fail. Use only in tests.
func Must[T any](result T, err error) T {
	if err != nil {
		panic(err)
	}
	return result
}

// Must2 panics if the second arg is an error.
//   - Allowed in production code only if it never fails.
//   - Don't use in production code if it is likely to fail. Use only in tests.
func Must2[T1, T2 any](result1 T1, result2 T2, err error) (T1, T2) {
	if err != nil {
		panic(err)
	}
	return result1, result2
}

// Must3 panics if the second arg is an error.
//   - Allowed in production code only if it never fails.
//   - Don't use in production code if it is likely to fail. Use only in tests.
func Must3[T1, T2, T3 any](result1 T1, result2 T2, result3 T3, err error) (T1, T2, T3) {
	if err != nil {
		panic(err)
	}
	return result1, result2, result3
}

// Must4 panics if the second arg is an error.
//   - Allowed in production code only if it never fails.
//   - Don't use in production code if it is likely to fail. Use only in tests.
func Must4[T1, T2, T3, T4 any](result1 T1, result2 T2, result3 T3, result4 T4, err error) (T1, T2, T3, T4) {
	if err != nil {
		panic(err)
	}
	return result1, result2, result3, result4
}

// Must5 panics if the second arg is an error.
//   - Allowed in production code only if it never fails.
//   - Don't use in production code if it is likely to fail. Use only in tests.
func Must5[T1, T2, T3, T4, T5 any](result1 T1, result2 T2, result3 T3, result4 T4, result5 T5, err error) (T1, T2, T3, T4, T5) {
	if err != nil {
		panic(err)
	}
	return result1, result2, result3, result4, result5
}
