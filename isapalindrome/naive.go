package isapalindrome

import "regexp"

func isPalindrome(s string) bool {
	r, _ := regexp.Compile("[^a-zA-Z\\d]")
	s = r.ReplaceAllString(s, "")
	// Edge case for empty string
	if len(s) == 0 {
		return true
	}

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s_i, s_j := s[i], s[j]
		if s[i] <= 'Z' {
			s_i += 32
		}
		if s[j] <= 'Z' {
			s_j += 32
		}
		if s_i != s_j {
			return false
		}
	}
	return true
}
