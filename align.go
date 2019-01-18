package align

import "fmt"

func Center(w int, s string) string {
	return fmt.Sprintf("%[1]*s", -w, fmt.Sprintf("%[1]*s", (w+len(s))/2, s))
}

func Right(w int, s string) string {
	return fmt.Sprintf("%[1]*s", w, s)
}

func Left(w int, s string) string {
	return fmt.Sprintf("%[1]*s", -w, s)
}
