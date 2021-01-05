package main

import "fmt"

func reverseVowels(s string) string {
	chars := []rune(s)
	for l, r := 0, len(chars)-1; l < r; {
		if !isVowel(chars[l]) {
			l++
			continue
		}
		if !isVowel(chars[r]) {
			r--
			continue
		}
		chars[l], chars[r] = chars[r], chars[l]
		l++
		r--
	}
	return string(chars)
}

func isVowel(char rune) bool {
	return char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' ||
		char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U'
}

func main() {
	test := "hello"
	fmt.Printf("%s", reverseVowels(test))

}
