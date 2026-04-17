package format

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/IamNanjo/go-logging"
)

const (
	wrapCutoff   = 70 // How many characters until text wraps on each line.
	prefixLength = 5  // Constant prefix length to keep each line aligned.
)

var (
	errPrefix     = []byte("\n  -> ") // Prefix for the error.
	errLinePrefix = []byte("\n     ") // Prefix for each wrapped line.
)

type multiError struct {
	message string  // Error message
	errs    []error // Nested errors
}

func (ne *multiError) Error() string {
	if ne == nil {
		return ""
	}
	var sb strings.Builder
	sb.WriteString(ne.message)
	last := len(ne.errs) - 1
	for i, err := range ne.errs {
		sb.WriteString(err.Error())
		if i < last {
			sb.Write([]byte(": "))
		}
	}
	return sb.String()
}

func (ne *multiError) Unwrap() []error {
	if ne == nil {
		return nil
	}
	return ne.errs
}

// Consistent error format with wrapping.
func Err(input string, args ...any) error {
	var errResult multiError

	var (
		textResult   strings.Builder
		lineLength   int
		argIndex     = 0
		lastWasColon = false
		in           = []rune(input)
		inLen        = len(in)
	)

	textResult.Grow(prefixLength + inLen)

	textResult.Write(errPrefix)
	lineLength = prefixLength

	wrap := func() {
		textResult.Write(errLinePrefix)
		lineLength = prefixLength
	}

	checkWrap := func(runes []rune) {
		for _, r := range runes {
			isSpace := unicode.IsSpace(r)
			// Wrap when word ends with ": ". Common in errors and this can improve readability.
			if lastWasColon && isSpace {
				wrap()
			} else if isSpace && lineLength > wrapCutoff { // Line is long enough to wrap
				wrap()
			} else {
				switch r {
				case '\t':
					textResult.Write([]byte("  "))
				default:
					textResult.WriteRune(r)
				}
				lineLength++
			}

			lastWasColon = r == ':'
		}
	}

	for i := 0; i < len(in); i++ {
		r := in[i]

		if r != '%' {
			checkWrap(in[i : i+1])
			continue
		}

		specifier, verb := parseSpecifier(in[i:])
		if verb == nil {
			continue
		}
		verbVal := *verb
		if verbVal == '%' {
			textResult.WriteByte('%')
			i++
			continue
		}

		if argIndex >= len(args) {
			logging.Default.Err("Not enough arguments in format.Err()\n")
			for _, sr := range specifier {
				textResult.WriteRune(sr)
			}
			i += len(specifier) - 1
		}
		arg := args[argIndex]
		argIndex++

		if verbVal == 'w' {
			if err, ok := arg.(error); ok {
				errResult.errs = append(errResult.errs, err)
				i++
				continue
			}
			*verb = 'v'
		}

		parsed := fmt.Sprintf(string(specifier), arg)
		i += len(specifier) - 1
		checkWrap([]rune(parsed))
	}

	errResult.message = textResult.String()

	return &errResult
}
