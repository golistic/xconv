// Copyright (c) 2022, Geert JM Vanderkelen

package xconv

import (
	"testing"

	"github.com/geertjanvdk/xkit/xt"
)

func TestUnsignedAsUint64(t *testing.T) {
	t.Run("valid argument type", func(t *testing.T) {
		exp := uint64(123)
		var cases = []any{
			uint(exp),
			uint8(exp),
			uint16(exp),
			uint32(exp),
			exp, // uint64
		}

		for _, c := range cases {
			t.Run("", func(t *testing.T) {
				xt.Eq(t, exp, UnsignedAsUint64(c))
			})
		}
	})

	t.Run("invalid argument type", func(t *testing.T) {
		var cases = []any{
			1, // int
			int8(1),
			int16(1),
			int32(1),
			int64(1),
			float32(1),
			float64(1),
			"string",
			[]byte("byte"),
		}

		for _, c := range cases {
			t.Run("", func(t *testing.T) {
				xt.Panics(t, func() {
					_ = UnsignedAsUint64(c)
				})
			})
		}
	})
}

func TestSignedAsUint64(t *testing.T) {
	t.Run("valid argument type", func(t *testing.T) {
		exp := int64(123)
		var cases = []any{
			int(exp),
			int8(exp),
			int16(exp),
			int32(exp),
			exp, // int64
		}

		for _, c := range cases {
			t.Run("", func(t *testing.T) {
				xt.Eq(t, exp, SignedAsInt64(c))
			})
		}
	})

	t.Run("invalid argument type", func(t *testing.T) {
		var cases = []any{
			uint(1),
			uint8(1),
			uint16(1),
			uint32(1),
			uint64(1),
			float32(1),
			float64(1),
			"string",
			[]byte("byte"),
		}

		for _, c := range cases {
			t.Run("", func(t *testing.T) {
				xt.Panics(t, func() {
					_ = SignedAsInt64(c)
				})
			})
		}
	})
}
