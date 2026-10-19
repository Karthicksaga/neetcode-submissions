func lengthOfLongestSubstring(s string) int {

	stringRune := []rune(s)

	charSet := make(map[rune]bool)
	l := 0
	maxLen := 0

	for r := 0; r < len(stringRune); r += 1{

		for charSet[stringRune[r]]{
			delete(charSet, stringRune[l])
			l += 1
		}

		charSet[stringRune[r]] = true

		if (r - l) + 1 > maxLen{
			maxLen = (r - l) + 1
		}

	}

	return maxLen

}
