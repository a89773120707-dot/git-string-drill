package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("drill started")

	fmt.Println(ToUpperFirst("hello"))
	fmt.Println(ToUpperFirst("anar"))

	fmt.Println(CountVowles("aeiouy"))

	fmt.Println(IsPalindrome("radar"))
	fmt.Println(IsPalindrome("race a car"))

	fmt.Println(ReverseWords("раз два  три    четыре"))
}

func ToUpperFirst(s string) string {
	if s == "" {
		return ""
	}

	return strings.ToUpper(string(s[0])) + s[1:]
}

func CountVowles(s string) int {
	count := 0
	lower := strings.ToLower(s)

	for _, r := range lower {
		switch r {
		case 'a', 'e', 'i', 'o', 'u', 'y':
			count++
		}
	}
	return count
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
