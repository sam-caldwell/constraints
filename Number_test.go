package types

import (
	"testing"
)

func NumberFunc[T Number](i T) bool {
	return any(i) == any(i)
}

func TestNumber(t *testing.T) {
	assert := func(c bool) {
		if !c {
			t.Fatal("error")
		}
	}

	assert(NumberFunc(int(0)))
	assert(NumberFunc(int8(0)))
	assert(NumberFunc(int16(0)))
	assert(NumberFunc(int32(0)))
	assert(NumberFunc(int64(0)))

	assert(NumberFunc(uint(0)))
	assert(NumberFunc(uint8(0)))
	assert(NumberFunc(uint16(0)))
	assert(NumberFunc(uint32(0)))
	assert(NumberFunc(uint64(0)))

	assert(NumberFunc(float32(0)))
	assert(NumberFunc(float64(0)))

}
