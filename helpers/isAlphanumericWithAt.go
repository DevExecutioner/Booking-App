package helpers

// Add this function anywhere in your code (after imports is fine)
func IsAlphanumericWithAt(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if !((r >= 'a' && r <= 'z') ||
			(r >= 'A' && r <= 'Z') ||
			(r >= '0' && r <= '9') ||
			r == '@') {
			return false
		}
	}
	return true
}
