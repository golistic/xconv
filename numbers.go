// Copyright (c) 2022, Geert JM Vanderkelen

package xconv

import "fmt"

// UnsignedAsUint64 returns n as uint64.
// Panics when n is not uint64, uint, uint8, uint16, uint32.
func UnsignedAsUint64(n any) uint64 {
	switch v := n.(type) {
	case uint64:
		return v
	case uint:
		return uint64(v)
	case uint8:
		return uint64(v)
	case uint16:
		return uint64(v)
	case uint32:
		return uint64(v)
	default:
		panic(fmt.Sprintf("n must be uint64, uint, uint8, uint16, uint32; not '%T'", n))
	}
}

// SignedAsInt64 returns n as int64.
// Panics when n is not int64, int, int8, int16, int32.
func SignedAsInt64(n any) int64 {
	switch v := n.(type) {
	case int64:
		return v
	case int:
		return int64(v)
	case int8:
		return int64(v)
	case int16:
		return int64(v)
	case int32:
		return int64(v)
	default:
		panic(fmt.Sprintf("n must be int64, int, int8, int16, int32; not '%T'", n))
	}
}
