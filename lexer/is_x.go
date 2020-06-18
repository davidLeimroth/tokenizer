package lexer

import "unicode"

func isLetter(r rune) bool {
	return unicode.IsLetter(r)
	// return 'a' <= r && r <= 'z' || 'A' <= r && r <= 'Z' || 192 <= r && r <= 214 || 216 <= r && r <= 222 || 223 <= r && r <= 246 || 248 <= r && r <= 255
}

func isDigit(r rune, nextRune rune) bool {
	// return unicode.IsDigit(r)
	return '0' <= r && r <= '9' || r == '.' && '0' <= nextRune && nextRune <= '9' || r == ',' && '0' <= nextRune && nextRune <= '9'
}
