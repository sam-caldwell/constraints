package types

// Number - Any native go numeric type (signed and unsigned integers, floats)
type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}
