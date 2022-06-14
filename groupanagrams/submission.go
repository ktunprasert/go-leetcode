package groupanagrams

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	groups := make(map[string]*[]string, 0)
	out := [][]string{}

	for i, str := range strs {
		key := []rune(str)
		sort.Slice(key, func(i, j int) bool { return key[i] < key[j] })
		strkey := string(key)
		switch {
		case i == 0, groups[strkey] == nil:
			groups[strkey] = &[]string{str}
		default:
			(*groups[strkey]) = append((*groups[strkey]), str)
		}
	}
	for _, v := range groups {
		out = append(out, *v)
	}
	return out
}
