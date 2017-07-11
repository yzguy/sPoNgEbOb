package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	input = os.Args[1]
)

func main() {
	fmt.Println(sPoNgEbOb(input))
}

func sPoNgEbOb(input string) string {
	for i, v := range input {
		s := string(v)
		if i%2 == 0 {
			output += strings.ToLower(s)
		} else {
			output += strings.ToUpper(s)
		}
	}

	return output
}
