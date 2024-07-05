package types

import (
	"math/big"
	"testing"
)

func testAnyNumberFunc[T AnyNumber](n T) bool {
	return true
}

func TestAnyNumber(t *testing.T) {
	assert := func(c bool) {
		if !c {
			t.Fatal("error")
		}
	}

	assert(testAnyNumberFunc(int(0)))
	assert(testAnyNumberFunc(int8(0)))
	assert(testAnyNumberFunc(int16(0)))
	assert(testAnyNumberFunc(int32(0)))
	assert(testAnyNumberFunc(int64(0)))

	assert(testAnyNumberFunc(uint(0)))
	assert(testAnyNumberFunc(uint8(0)))
	assert(testAnyNumberFunc(uint16(0)))
	assert(testAnyNumberFunc(uint32(0)))
	assert(testAnyNumberFunc(uint64(0)))

	assert(testAnyNumberFunc(float32(0)))
	assert(testAnyNumberFunc(float64(0)))

	assert(testAnyNumberFunc(big.Int{}))
	assert(testAnyNumberFunc(big.Float{}))

}
