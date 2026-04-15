package format

import "unicode"

// parseSpecifier returns the specifier slice and a pointer to the verb within that slice.
func parseSpecifier(input []rune) (specifier []rune, verb *rune) {
	if len(input) < 2 || input[0] != '%' {
		return nil, nil
	}

	i := 1

	// Handle literal percent %%
	if input[i] == '%' {
		specifier = input[0:2]
		verb = &input[1] // Pointer to the second '%'
		return
	}

	// Track start of flags
	for i < len(input) && isFlag(input[i]) {
		i++
	}

	// Skip Width
	i = skipDigitsOrAsterisk(input, i)

	// Skip Precision
	if i < len(input) && input[i] == '.' {
		i++
		i = skipDigitsOrAsterisk(input, i)
	}

	// Capture the verb as a pointer to the specific element in the slice
	if i < len(input) {
		specifier = input[0 : i+1]
		verb = &input[i]
	}

	return
}

func isFlag(r rune) bool {
	switch r {
	case '+', '-', '#', ' ', '0':
		return true
	}
	return false
}

func skipDigitsOrAsterisk(input []rune, i int) int {
	if i < len(input) && input[i] == '*' {
		return i + 1
	}
	for i < len(input) && unicode.IsDigit(input[i]) {
		i++
	}
	return i
}
