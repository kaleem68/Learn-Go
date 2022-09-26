package main

import "unicode"

// Letter case is ignored, as are non-letters.
func IsPalindrome(s string) bool {
	var letters []rune
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}

// IsPalindrome’s use of byte sequences, not rune sequences: so "été" fails
// func IsPalindrome(s string) bool {
// 	for i := range s {
// 		if s[i] != s[len(s)-1-i] {
// 			return false
// 		}
// 	}
// 	return true
// }
