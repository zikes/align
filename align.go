package align

import (
	"fmt"
	"unicode/utf8"
)

// Center will center text within a given length of space
func Center(w int, s string) string {
	return fmt.Sprintf("%[1]*s", -w, fmt.Sprintf("%[1]*s", (w+utf8.RuneCountInString(s))/2, s))
}

// Right will right-align text within a given length of space
func Right(w int, s string) string {
	return fmt.Sprintf("%[1]*s", w, s)
}

// Left will left-align text within a given length of space
func Left(w int, s string) string {
	return fmt.Sprintf("%[1]*s", -w, s)
}
