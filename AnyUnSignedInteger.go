package types

// AnyUnsignedInteger - Any native go numeric unsigned integer type (signed and unsigned integers)
type AnyUnsignedInteger interface {
	uint | uint8 | uint16 | uint32 | uint64
}
