package groupanagrams

func isAnagram(s *string, t *string) bool {
	slen := len(*s)
	tlen := len(*t)
	if slen != tlen {
		return false
	}

	alphas := make(map[byte]int8, 26)
	for i := 0; i < slen; i++ {
		alphas[(*s)[i]]++
		alphas[(*t)[i]]--
	}
	for _, v := range alphas {
		if v != 0 {
			return false
		}
	}

	return true
}

func groupAnagramsNaive(strs []string) [][]string {
	groups := make(map[string][]string, 0)
	out := [][]string{}
	for _, str := range strs {
		switch {
		case len(groups) == 0:
			groups[str] = []string{str}
		default:
			found := false
			for key, arr := range groups {
				if len(str) == len(key) && isAnagram(&key, &str) {
					groups[key] = append(arr, str)
					found = true
				}
			}
			if !found {
				groups[str] = []string{str}
			}
		}
	}
	for _, v := range groups {
		out = append(out, v)
	}
	return out
}
