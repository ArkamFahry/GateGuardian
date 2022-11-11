package utils

// StringSliceContains checks if a string slice contains a particular string
func StringSliceContains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
