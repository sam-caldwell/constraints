package types

import (
	"testing"
)

func AnySignedIntegerFunc[T AnySignedInteger](i T) bool {
	return any(i) == any(i)
}

func TestAnySignedInteger(t *testing.T) {
	assert := func(c bool) {
		if !c {
			t.Fatal("error")
		}
	}

	assert(AnySignedIntegerFunc(int(0)))
	assert(AnySignedIntegerFunc(int8(0)))
	assert(AnySignedIntegerFunc(int16(0)))
	assert(AnySignedIntegerFunc(int32(0)))
	assert(AnySignedIntegerFunc(int64(0)))
}
