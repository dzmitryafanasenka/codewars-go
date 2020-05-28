package kata

import (
	"bytes"
	"strings"
)

func DuplicateEncode(word string) string {
	frequency := make(map[rune]int)
	word = strings.ToLower(word)
	buf := bytes.NewBufferString("")

	for _, char := range word {
		count, _ := frequency[char]
		frequency[char] = count + 1
	}

	for _, char := range word {
		count, _ := frequency[char]
		if count > 1 {
			buf.WriteString(")")
		} else {
			buf.WriteString("(")
		}
	}

	return buf.String()
}
