package strutil

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	runes := []rune(s)
	len := len(runes)
	for i := 0; i < len/2; i++ {
		runes[i], runes[len-1-i] = runes[len-1-i], runes[i]
	}
	return string(runes)
}
