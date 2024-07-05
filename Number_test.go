package types

import (
	"testing"
)

func testNumberFunc[T Number](n T) bool {
	return n == 0
}

func TestNumber(t *testing.T) {
	assert := func(c bool) {
		if !c {
			t.Fatal("error")
		}
	}

	assert(testNumberFunc(int(0)))
	assert(testNumberFunc(int8(0)))
	assert(testNumberFunc(int16(0)))
	assert(testNumberFunc(int32(0)))
	assert(testNumberFunc(int64(0)))

	assert(testNumberFunc(uint(0)))
	assert(testNumberFunc(uint8(0)))
	assert(testNumberFunc(uint16(0)))
	assert(testNumberFunc(uint32(0)))
	assert(testNumberFunc(uint64(0)))

	assert(testNumberFunc(float32(0)))
	assert(testNumberFunc(float64(0)))

}
