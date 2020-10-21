package main

import (
	"fmt"
	"strings"
)

func CaseSensitive(s, t string) string {
	if s == t {
		return "same"
	} else if strings.ToLower(s) == strings.ToLower(t) {
		return "case-insensitive"
	} else {
		return "different"
	}
}

func main() {
	var s, t string
	fmt.Scan(&s, &t)

	fmt.Println(CaseSensitive(s, t))
}
