package funcs

import (
	"strconv"
	"strings"
)

// (cap)
func Ucl(words []string) []string {
	for i := 0; i < len(words); i++ {
		val := words[i]

		// (cap)
		if strings.HasPrefix(val, "(cap") {
			if val == "(cap)" && i > 0 {
				Cap(words, 1, i-1)
				words = append(words[:i], words[i+1:]...)
				i--
			} else if strings.HasSuffix(val, ")") {
				k := TakeNumFromString(val)
				Cap(words, k, i-1)
				words = append(words[:i], words[i+1:]...)
				i--
			} else if strings.HasPrefix(val, "(cap,") && i > 0 {
				k := TakeNumFromString(val)
				Cap(words, k, i-1)
				words = append(words[:i], words[i+1:]...)
				i--
			}
		}

		// (low)
		if strings.HasPrefix(val, "(low") {
			if val == "(low)" && i > 0 {
				Low(words, 1, i-1)
				words = append(words[:i], words[i+1:]...)
				i--
			} else if strings.HasSuffix(val, ")") {
				k := TakeNumFromString(val)
				Low(words, k, i-1)
				words = append(words[:i], words[i+1:]...)
				i--
			}
		}

		// (up)
		if strings.HasPrefix(val, "(up") {
			if val == "(up)" && i > 0 {
				Up(words, 1, i-1)
				words = append(words[:i], words[i+1:]...)
				i--
			} else if strings.HasSuffix(val, ")") {
				k := TakeNumFromString(val)
				Up(words, k, i-1)
				words = append(words[:i], words[i+1:]...)
				i--
			}
		}
	}

	return words

}

func Cap(s []string, n int, m int) {
	for n > 0 && m >= 0 {
		if IsWord(s[m]) {
			tp := FindFl(s[m])
			if len(s[m]) == 1 {
				s[m] = strings.ToUpper(s[m])
			} else {
				s[m] = strings.ToUpper(s[m][:tp+1]) + strings.ToLower(s[m][tp+1:])
			}
			n--
		}
		m--
	}
}

func Low(s []string, n int, m int) {
	for n > 0 && m >= 0 {
		if IsWord(s[m]) {
			s[m] = strings.ToLower(s[m])
			n--
		}
		m--
	}
}

func Up(s []string, n int, m int) {
	for n > 0 && m >= 0 {
		if IsWord(s[m]) {
			s[m] = strings.ToUpper(s[m])
			n--
		}
		m--
	}
}

func TakeNumFromString(s string) int {
	res := ""
	for _, val := range s {
		if val >= '0' && val <= '9' {
			res += string(val)
		}
	}
	if res == "" {
		return 0
	}
	ans, _ := strconv.Atoi(res)
	return ans
}

func FindFl(s string) int {
	for i, val := range s {

		if val >= 'a' && val <= 'z' || val >= 'A' && val <= 'Z' {
			return i
		}
	}
	return -1
}
