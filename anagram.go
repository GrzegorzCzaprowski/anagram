package anagram

func Anagram(s string) string {
	runes := []rune(s)
	runes2 := []rune(s)
	for i, r := range runes {
		runes2[len(s)-i-1] = r
	}

	return string(runes2)
}
