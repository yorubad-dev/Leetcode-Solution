package main

import "strings"

func reverseWords(s string) string {
	words := strings.Fields(s)
	for left, right := 0, len(words)-1; left < right; left, right = left+1, right-1 {
		words[left], words[right] = words[right], words[left]
	}

	return strings.Join(words, " ")
}

func main() {
	reverseWords(" Hello World ")
}
