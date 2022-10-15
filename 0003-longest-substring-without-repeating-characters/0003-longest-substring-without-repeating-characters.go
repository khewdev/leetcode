func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	lo, hi, ans := 0, 0, 1.0
	m := make(map[byte]bool)

	for hi < len(s) {
		c := s[hi]
		hi++
		for lo < hi {
			if _, present := m[c]; !present {
				break
			}
			lowC := s[lo]
			lo++
			delete(m, lowC)
		}
		m[c] = true
		ans = math.Max(ans, float64(hi-lo))
	}
	return int(ans)
}