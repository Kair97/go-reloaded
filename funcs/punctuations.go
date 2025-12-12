package funcs

import "strings"

const alp = "abcdefghijklmnopqrstuvwxyz"

// Separate leading/trailing punctuation from words, including quotes
func SeparatePunc(words []string) []string {
	var res []string
	for _, w := range words {
		prefix := ""
		suffix := ""
		for len(w) > 0 && strings.Contains(".,;:!?'\"", string(w[0])) { // added quotes
			prefix += string(w[0])
			w = w[1:]
		}
		for len(w) > 0 && strings.Contains(".,;:!?'\"", string(w[len(w)-1])) { // added quotes
			suffix = string(w[len(w)-1]) + suffix
			w = w[:len(w)-1]
		}
		if prefix != "" {
			res = append(res, prefix)
		}
		if w != "" {
			res = append(res, w)
		}
		if suffix != "" {
			res = append(res, suffix)
		}
	}
	return res
}

// Reattach punctuation after capitalization, including quotes
func ReattachPunc(words []string) []string {
	var res []string
	for i := 0; i < len(words); i++ {
		w := words[i]
		// only attach punctuation that is not a word
		if !IsWord(w) && len(res) > 0 {
			res[len(res)-1] += w
		} else {
			res = append(res, w)
		}
	}
	return res
}

func IsWord(s string) bool {
	for _, val := range s {
		if strings.Contains(alp, strings.ToLower(string(val))) {
			return true
		}
	}
	return false
}
