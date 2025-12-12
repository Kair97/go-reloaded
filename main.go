package main

import (
	"fmt"
	"go-reloaded/funcs"
	"os"
	"strings"
)

const alp = "abcdefghijklmnopqrstuvwxyz"

func main() {

	fmt.Println()
	if len(os.Args) != 3 {
		fmt.Println("Error ! ")
		return
	}
	readF := os.Args[1]
	writeF := os.Args[2]

	content, _ := os.ReadFile(readF)

	// Handle multiple lines correctly
	lines := strings.Split(string(content), "\n")
	for li, line := range lines {
		words := strings.Fields(line)

		// Preprocess punctuation: separate leading/trailing punctuation
		words = funcs.SeparatePunc(words)

		// Markup processing: (cap), (low), (up)
		for i := 0; i < len(words); i++ {

			words = funcs.Ucl(words)

			words = funcs.ReattachPunc(words)
			words = funcs.MergeQuotes(words)
			words = funcs.MergeDQuotes(words)
			words = funcs.FixArticles(words)
			lines[li] = strings.Join(words, " ")
		}
	}

	contPaste := strings.Join(lines, "\n")

	os.WriteFile(writeF, []byte(contPaste), 0o644)

	contR, _ := os.ReadFile(writeF)
	fmt.Printf("Initially: %v\n", string(content))
	fmt.Println()
	fmt.Printf("Result: %v\n", string(contR))
	fmt.Println()
}
