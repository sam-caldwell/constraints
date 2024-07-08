package types

import (
	"testing"
)

func AnyUnSignedIntegerFunc[T AnyUnsignedInteger](i T) bool {
	return any(i) == any(i)
}

func TestAnyUnSignedInteger(t *testing.T) {
	assert := func(c bool) {
		if !c {
			t.Fatal("error")
		}
	}

	assert(AnyUnSignedIntegerFunc(uint(0)))
	assert(AnyUnSignedIntegerFunc(uint8(0)))
	assert(AnyUnSignedIntegerFunc(uint16(0)))
	assert(AnyUnSignedIntegerFunc(uint32(0)))
	assert(AnyUnSignedIntegerFunc(uint64(0)))
}
