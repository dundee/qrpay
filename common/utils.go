package common

// TrimToLength shortends given string to given length
func TrimToLength(value string, length int) string {
	if len(value) > length {
		return value[:length]
	}
	return value
}
