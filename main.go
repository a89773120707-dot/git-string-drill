package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("drill started")

	fmt.Println(TrimAndUpper("   hello world   "))
	fmt.Println(TrimAndUpper("   git init   "))

	fmt.Println(ToUpperFirst("hello"))
	fmt.Println(ToUpperFirst("anar"))
}

func TrimAndUpper(s string) string {
	sTrim := strings.TrimSpace(strings.ToUpper(s))
	return sTrim
}

func ToUpperFirst(s string) string {
	if s == "" {
		return ""
	}

	return strings.ToUpper(string(s[0])) + s[1:]
}
