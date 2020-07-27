// Package umath provides some utilities for the uint type
// that mirror those found in the built-in math package.
package umath

// Code generated by generate/generate.go from math.go.tmpl. DO NOT EDIT.

// Dim returns the maximum of x-y or 0.
func Dim(x, y uint) uint {
	if y >= x {
		return 0
	}

	return x - y

}

// Max returns the larger of x or y.
func Max(x, y uint) uint {
	if x > y {
		return x
	}
	return y
}

// Min returns the smaller of x or y.
func Min(x, y uint) uint {
	if x < y {
		return x
	}
	return y
}