package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(text string) (string, error) {
	var sb strings.Builder
	if len(text) > 0 {
		for i := 0; i < len(text); i++ {
			var prev, curr, next rune
			if i > 0 {
				prev = rune(text[i-1])
			}
			curr = rune(text[i])
			if i < len(text)-1 {
				next = rune(text[i+1])
			}

			err := handleRune(&sb, prev, curr, next)
			if err != nil {
				return "", err
			}
		}
	} else {
		return "", nil
	}
	return sb.String(), nil
}

func handleRune(sb *strings.Builder, prev, curr, next rune) error {
	if prev == 0 {
		if unicode.IsDigit(curr) {
			return ErrInvalidString
		}
	}
	if unicode.IsDigit(prev) && unicode.IsDigit(curr) {
		return ErrInvalidString
	}

	if prev == '\\' {
		if !unicode.IsDigit(curr) && curr != '\\' {
			return ErrInvalidString
		}
	} else if unicode.IsDigit(curr) || curr == '\\' {
		return nil
	}
	if unicode.IsDigit(next) && next != 0 {
		rptCount, err := strconv.Atoi(string(next))
		if err != nil {
			return ErrInvalidString
		}
		repeatedLetters := strings.Repeat(string(curr), rptCount)
		sb.WriteString(repeatedLetters)
	} else {
		sb.WriteRune(curr)
	}
	return nil
}
