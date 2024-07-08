package types

import (
	"testing"
)

func AnyIntegerFunc[T AnyInteger](i T) bool {
	return any(i) == any(i)
}

func TestAnyInteger(t *testing.T) {
	assert := func(c bool) {
		if !c {
			t.Fatal("error")
		}
	}

	assert(AnyIntegerFunc(int(0)))
	assert(AnyIntegerFunc(int8(0)))
	assert(AnyIntegerFunc(int16(0)))
	assert(AnyIntegerFunc(int32(0)))
	assert(AnyIntegerFunc(int64(0)))

	assert(AnyIntegerFunc(uint(0)))
	assert(AnyIntegerFunc(uint8(0)))
	assert(AnyIntegerFunc(uint16(0)))
	assert(AnyIntegerFunc(uint32(0)))
	assert(AnyIntegerFunc(uint64(0)))
}
