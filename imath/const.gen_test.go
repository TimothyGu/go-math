package imath

// Code generated by generate/generate.go from const_test.go.tmpl. DO NOT EDIT.

import "testing"

func TestMaxUntyped(t *testing.T) {

	t.Run("int", func(t *testing.T) {
		var _ int = MaxValue
	})

	t.Run("int64", func(t *testing.T) {
		var _ int64 = MaxValue
	})

	t.Run("uint64", func(t *testing.T) {
		var _ uint64 = MaxValue
	})

	t.Run("uint", func(t *testing.T) {
		var _ uint = MaxValue
	})

}

func TestMinUntyped(t *testing.T) {

	t.Run("int", func(t *testing.T) {
		var _ int = MinValue
	})

	t.Run("int64", func(t *testing.T) {
		var _ int64 = MinValue
	})

}