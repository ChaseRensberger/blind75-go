func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	n := len(s)
	cs, ct := make(map[rune]int), make(map[rune]int)
	for i := 0; i < n; i++ {
		cs[rune(s[i])]++
		ct[rune(t[i])]++
	}
	for k, v := range cs {
		if (ct[k]) != v {
			return false
		}
	}
	return true
}
