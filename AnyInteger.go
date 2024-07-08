package types

// AllInteger - Any native go numeric signed or unsigned integer type (signed and unsigned integers)
type AllInteger interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}
