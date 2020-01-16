package main

import (
	"fmt"
	"strings"
)

func main() {
	var word string
	fmt.Print("Enter your word=")
	fmt.Scanf("%s", &word)
	if strings.HasPrefix(strings.ToLower(word), strings.ToLower("i")) &&
	    strings.HasSuffix(strings.ToLower(word), strings.ToLower("n")) &&
	    strings.Contains(strings.ToLower(word), strings.ToLower("a")) {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
