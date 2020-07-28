package umath

// Code generated by generate/generate.go from const_test.go.tmpl. DO NOT EDIT.

import "testing"

func TestMaxUntyped(t *testing.T) {

	t.Run("uint", func(t *testing.T) {
		var _ uint = MaxValue
	})

	t.Run("uint64", func(t *testing.T) {
		var _ uint64 = MaxValue
	})

}

func TestMinUntyped(t *testing.T) {

	t.Run("int64", func(t *testing.T) {
		var _ int64 = MinValue
	})

	t.Run("int32", func(t *testing.T) {
		var _ int32 = MinValue
	})

	t.Run("int16", func(t *testing.T) {
		var _ int16 = MinValue
	})

	t.Run("int8", func(t *testing.T) {
		var _ int8 = MinValue
	})

	t.Run("int", func(t *testing.T) {
		var _ int = MinValue
	})

	t.Run("uint64", func(t *testing.T) {
		var _ uint64 = MinValue
	})

	t.Run("uint32", func(t *testing.T) {
		var _ uint32 = MinValue
	})

	t.Run("uint16", func(t *testing.T) {
		var _ uint16 = MinValue
	})

	t.Run("uint8", func(t *testing.T) {
		var _ uint8 = MinValue
	})

	t.Run("uint", func(t *testing.T) {
		var _ uint = MinValue
	})

}