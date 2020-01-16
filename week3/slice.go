package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var word string
	var slice = make([]int, 3)

	for i := 0; i < len(slice); i++ {
		fmt.Print("Enter your number=")
		fmt.Scanf("%s", &word)
		if strings.EqualFold(strings.ToUpper(word), "X") {
			break
		}
		numberToAdd, _ := strconv.Atoi(word)
		slice = append(slice, numberToAdd)
		sort.Sort(sort.IntSlice(slice))
		fmt.Println(slice)
	}
}
