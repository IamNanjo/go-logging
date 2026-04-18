package formatting

// Add spaces until input is at least as long as size
func GetPadding(length int, targetLength int) []byte {
	padding := targetLength - length
	if padding <= 0 {
		return []byte{}
	}

	result := make([]byte, padding)
	for i := range padding {
		result[i] = ' '
	}

	return result
}
