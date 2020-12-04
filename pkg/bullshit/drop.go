package bullshit

// DropErrorBoolean drops the error and returns with the boolean value only.
func DropErrorBoolean(value bool, err error) bool {
	return value
}

// DropErrorInt64 drops the error and returns with the Int64 value only.
func DropErrorInt64(value int64, err error) int64 {
	return value
}
