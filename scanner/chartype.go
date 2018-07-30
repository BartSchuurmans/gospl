package scanner

func isAlpha(c byte) bool {
	return c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z'
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isWord(c byte) bool {
	return isAlpha(c) || isDigit(c) || c == '_'
}
