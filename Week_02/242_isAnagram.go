package Week_02

//判断有效的字母异位词
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	a := make(map[byte]int, 0)
	for i := range s {
		a[s[i]]++
	}
	for i := range t {
		a[t[i]]--
		if a[t[i]] < 0 {
			return false
		}
	}
	return true
}
