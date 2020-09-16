package util

// BoolToInt true -> 1, false -> 0
func BoolToInt(b bool) int64 {
	if b {
		return 1
	}
	return 0
}

// MaxInt64 max of two numbers
func MaxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

// MinInt64 min of two numbers
func MinInt64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
