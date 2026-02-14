package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("drill started")

	fmt.Println(ToUpperFirst("hello"))
	fmt.Println(ToUpperFirst("anar"))

	fmt.Println(IsPalindrome("radar"))
	fmt.Println(IsPalindrome("race a car"))

	fmt.Println(ReverseWords("раз два  три    четыре"))

	fmt.Println(IsPalindromeUnicode("привет "))
	fmt.Println(IsPalindromeUnicode("прирп"))
}

func ToUpperFirst(s string) string {
	if s == "" {
		return ""
	}

	return strings.ToUpper(string(s[0])) + s[1:]
}

func IsPalindrome(s string) bool {
	cleaned := ""
	for _, r := range s {
		if r != ' ' {
			cleaned += string(r)
		}
	}

	left, right := 0, len(cleaned)-1

	for left < right {
		if cleaned[left] != cleaned[right] {
			return false
		}
		left++
		right--
	}
	return true

}

// [ "Раз", "Два", "Три", "Четыре" ]
func ReverseWords(s string) string {
	if s == "" {
		return ""
	}
	words := strings.Fields(s)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ")
}

func IsPalindromeUnicode(s string) bool {
	var runes []rune

	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			lower := unicode.ToLower(r)
			runes = append(runes, lower)
		}
	}

	left, right := 0, len(runes)-1
	for left < right {
		if runes[left] != runes[right] {
			return false
		}

		runes[left], runes[right] = runes[right], runes[left]

		left++
		right--
	}

	return true

}
