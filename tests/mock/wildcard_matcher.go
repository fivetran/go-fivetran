package mock

// WildcardMatch matches the specified string against the supplied pattern
// with wildcard characters ('*' and '?'). It uses the recursive algorithm
// which is not very efficient, its complexity can be up to O(N^k), where
// k = (number of asterisks in the pattern + 1). But it could be slightly
// faster than the matrix algorithm when the pattern contains just one asterisk.
func WildcardMatch(str string, pattern string) bool {
	strPos := 0
	strLen := len(str)
	for patPos := range pattern {
		if strPos == strLen {
			return false
		}

		patChar := pattern[patPos]

		if patChar == '*' {
			patternLeft := pattern[patPos+1:]
			for i := strPos; i <= strLen; i++ {
				if WildcardMatch(str[i:], patternLeft) {
					return true
				}
			}
			return false
		}

		if patChar != str[strPos] && patChar != '?' {
			return false
		}

		strPos++
	}

	return strPos == strLen
}
