package services

// works only with int64 values.
func Int64Abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}
