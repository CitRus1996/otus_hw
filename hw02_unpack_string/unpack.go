package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

type State int64

const (
	None State = iota
	Print
	DontPrint
	Screen
)

func Unpack(text string) (string, error) {
	var sb strings.Builder
	var state State
	var prev rune
	if len(text) == 0 {
		return "", nil
	}
	for _, curr := range text {
		if curr == '\\' {
			if state != Screen {
				state = Screen
			} else {
				state = Print
				sb.WriteRune(curr)
			}
			prev = curr
			continue
		}

		if unicode.IsLetter(curr) {
			state = Print
			sb.WriteRune(curr)
			prev = curr
			continue
		}

		if err := handleDigit(&sb, prev, curr, &state); err != nil {
			return "", err
		}
		prev = curr
	}

	return sb.String(), nil
}

func handleDigit(sb *strings.Builder, prev, curr rune, state *State) error {
	if unicode.IsDigit(curr) {
		switch *state {
		case Screen:
			*state = Print
			sb.WriteRune(curr)
		case Print:
			rptCount, err := strconv.Atoi(string(curr))
			if err != nil {
				return ErrInvalidString
			}
			if rptCount > 0 {
				repeated := strings.Repeat(string(prev), rptCount-1)
				sb.WriteString(repeated)
			} else {
				temp := sb.String()
				cut := temp[:len(temp)-1]
				sb.Reset()
				sb.WriteString(cut)
			}
			*state = DontPrint
		case None, DontPrint:
			return ErrInvalidString
		}
	}
	return nil
}
