package validanagram

func isAnagram(s string, t string) bool {
	slen := len(s)
	tlen := len(t)
	if slen != tlen {
		return false
	}

	alphas := make(map[byte]int8, 26)
	for i := 0; i < slen; i++ {
		alphas[s[i]]++
		alphas[t[i]]--
	}
	for _, v := range alphas {
		if v != 0 {
			return false
		}
	}

	return true
}
