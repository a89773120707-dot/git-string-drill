package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("drill started")

	fmt.Println(TrimAndUpper("   hello world   "))
	fmt.Println(TrimAndUpper("   git init   "))
}

func TrimAndUpper(s string) string {
	sTrim := strings.TrimSpace(strings.ToUpper(s))
	return sTrim
}
