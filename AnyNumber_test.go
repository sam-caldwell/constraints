package types

import (
	"math/big"
	"testing"
)

func AnyNumberFunc[T AnyNumber](i T) bool {
	return any(i) != nil
}

func TestAnyNumber(t *testing.T) {
	assert := func(c bool) {
		if !c {
			t.Fatal("error")
		}
	}

	assert(AnyNumberFunc(int(0)))
	assert(AnyNumberFunc(int8(0)))
	assert(AnyNumberFunc(int16(0)))
	assert(AnyNumberFunc(int32(0)))
	assert(AnyNumberFunc(int64(0)))

	assert(AnyNumberFunc(uint(0)))
	assert(AnyNumberFunc(uint8(0)))
	assert(AnyNumberFunc(uint16(0)))
	assert(AnyNumberFunc(uint32(0)))
	assert(AnyNumberFunc(uint64(0)))

	assert(AnyNumberFunc(float32(0)))
	assert(AnyNumberFunc(float64(0)))

	assert(AnyNumberFunc(big.Int{}))
	assert(AnyNumberFunc(big.Float{}))

}
