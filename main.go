package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hi")
	fmt.Println()
	readF := os.Args[1]
	writeF := os.Args[2]

	content, _ := os.ReadFile(readF)

	words := strings.Split(string(content), " ")
	// ss := words[1]
	// fmt.Println(ss[:5])

	for i := 0; i < len(words); i++ {
		val := words[i]

		if val == "," || val == ";" || val == "!" || val == "?" || val == ":" || val == "..." {
			if i != 0 {
				words[i-1] += val
				words = append(words[:i], words[i+1:]...)
				i--
			}
		}
	}

	// to CAP and LOW and UP :
	for i := 0; i < len(words); i++ {
		val := words[i]

		if len(val) >= 5 && val[:4] == "(cap" {
			if i == 0 {
				continue
			} else if len(val) == 5 && val == "(cap)" {
				Cap(words, 1, i-1)
			} else {
				k := TakeNumFromString(words[i+1])
				Cap(words, k, i-1)
				i++
			}
		} else if len(val) >= 5 && val[:4] == "(low" {
			if i == 0 {
				continue
			} else if len(val) == 5 && val == "(low)" {
				Low(words, 1, i-1)
			} else {
				k := TakeNumFromString(words[i+1])
				Low(words, k, i-1)
				fmt.Println(k)
				i++
			}
		} else if len(val) >= 4 && val[:3] == "(up" {
			if i == 0 {
				continue
			} else if len(val) == 4 && val == "(up)" {
				Up(words, 1, i-1)
			} else {
				k := TakeNumFromString(words[i+1])
				Up(words, k, i-1)
				i++
			}
		}

	}

	contPaste := strings.Join(words, " ")
	os.WriteFile(writeF, []byte(contPaste), 0644)

	contR, _ := os.ReadFile(writeF)
	fmt.Println(string(contR))

	fmt.Println()
}

func Cap(s []string, n int, m int) {
	for n > 0 && m >= 0 {
		if len(s[m]) == 1 {
			s[m] = strings.ToUpper(s[m])
		} else {
			s[m] = strings.ToUpper(s[m][:1]) + s[m][1:]
		}
		m--
		n--
	}
}

func Low(s []string, n int, m int) {
	for n > 0 && m >= 0 {
		s[m] = strings.ToLower(s[m])
		m--
		n--
	}
}

func Up(s []string, n int, m int) {
	for n > 0 && m >= 0 {
		s[m] = strings.ToUpper(s[m])
		m--
		n--
	}
}

func TakeNumFromString(s string) int {
	res := ""

	for _, val := range s {
		if val >= '0' && val <= '9' {
			res += string(val)
		} else {
			break
		}
	}
	ans, _ := strconv.Atoi(res)
	return ans
}
