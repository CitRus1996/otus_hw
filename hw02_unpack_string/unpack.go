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
	for i := 0; i < len(text); i++ {
		currentRune := rune(text[i])
		if unicode.IsDigit(currentRune) {
			return "", errors.New("некорректная строка")
		}
		var nextRune rune
		if i < len(text)-1 {
			nextRune = rune(text[i+1])
			if currentRune == '\\' {
				if unicode.IsDigit(nextRune) || nextRune == '\\' {
					currentRune = nextRune
					i++
					if i < len(text)-1 {
						nextRune = rune(text[i+1])
					} else {
						sb.WriteRune(currentRune)
						continue
					}
				} else {
					return "", errors.New("некорректная строка")
				}
			}
			if unicode.IsDigit(nextRune) {
				repeatCount, err := strconv.Atoi(string(nextRune))
				if err != nil {
					return "", err
				}
				rpt := strings.Repeat(string(currentRune), repeatCount)
				sb.WriteString(rpt)
				i++
			} else {
				sb.WriteRune(currentRune)
			}
		} else {
			sb.WriteRune(currentRune)
		}
	}
	return sb.String(), nil
	return "", nil
}
