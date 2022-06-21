package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	//Place your code here.
	fmt.Println(1, Unpack("a4bc2d5e"))
	fmt.Println(2, Unpack("abcd"))
	fmt.Println(3, Unpack("3abc"))
	fmt.Println(4, Unpack("45"))
	fmt.Println(5, Unpack("aaa10b"))
	fmt.Println(6, Unpack("aaa0be"))
	fmt.Println(7, Unpack(""))
	fmt.Println(8, Unpack("d\n5abc"))
	fmt.Println(9, Unpack(`qwe\4\5`))
	fmt.Println(10, Unpack(`qwe\45`))
	fmt.Println(11, Unpack(`qwe\\5`))
	fmt.Println(12, Unpack(`qw\ne`))
}

func Unpack(text string) string {
	var sb strings.Builder
	for i := 0; i < len(text); i++ {
		currentRune := rune(text[i])
		if unicode.IsDigit(currentRune) {
			return ""
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
					return ""
				}
			}
			if unicode.IsDigit(nextRune) {
				repeatCount, err := strconv.Atoi(string(nextRune))
				if err != nil {
					return ""
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
	return sb.String()
}
