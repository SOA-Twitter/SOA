package data

import (
	"fmt"
	"unicode"
)

func ValidatePassword(s string) error {
	var (
		sevenOrMore = false
		number      = false
		upper       = false
		lower       = false
		special     = false
	)
	letters := 0
	for _, c := range s {
		switch {
		case unicode.IsNumber(c):
			number = true
			letters++
		case unicode.IsUpper(c):
			upper = true
			letters++
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
			letters++
		case unicode.IsLetter(c) || c == ' ':
			lower = true
			letters++
		default:
		}
	}
	sevenOrMore = letters >= 7
	if sevenOrMore && number && upper && special && lower == true {
		return nil
	}
	return fmt.Errorf("pattern error")
}
