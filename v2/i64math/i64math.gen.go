// Package i64math provides some utilities for the int64 type
// that mirror those found in the built-in math package.
package i64math

// Code generated by generate/generate.go from math.go.tmpl. DO NOT EDIT.

// Abs returns the absolute value of x, or |x|. If |x| exceeds MaxValue,
// MaxValue is returned.
func Abs(x int64) int64 {
	if x < 0 {
		if -x < 0 {
			// x == MinValue
			return MaxValue
		}
		return -x
	}
	return x
}

// Dim returns the maximum of x-y or 0.
func Dim(x, y int64) uint64 {
	if y >= x {
		return 0
	}

	// Go guarantees 2's complement, so the conversion is well-defined.
	return uint64(x - y)

}

// Max returns the larger of x or y.
func Max(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

// Min returns the smaller of x or y.
func Min(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}
