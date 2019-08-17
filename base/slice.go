package base

// GetBetween - return the first item different of empty
func GetBetween(items []string) string {
	for _, item := range items {
		if item != "" {
			return item
		}
	}
	return ""
}

// ContainStr - return true if s contain e
func ContainStr(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
