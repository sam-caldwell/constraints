package types

// AnySignedInteger - Any native go numeric signed integer type (signed and unsigned integers)
type AnySignedInteger interface {
	int | int8 | int16 | int32 | int64
}
