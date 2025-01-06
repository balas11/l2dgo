package main

import (
	"fmt"
	"strings"
)

func updateCount(vc map[rune]int, c rune) {
	if _, found := vc[c]; found {
		vc[c] = vc[c] + 1
		// } else {
		// 	vc[c] = 1
	}
}

func main() {
	vowels := "aeiou"
	// vowelCount := make(map[rune]int)
	vowelCount := map[rune]int{
		'a': 0, 'e': 0, 'i': 0, 'o': 0, 'u': 0,
	}

	phrase := "the quick brown fox jumps over the lazy dog"

	for _, char := range phrase {
		if strings.ContainsRune(vowels, char) {
			updateCount(vowelCount, char)
		}
	}
	for key, val := range vowelCount {
		fmt.Printf("%c : %d\n", key, val)
	}
}
