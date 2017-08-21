package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) > 1 {
		input := strings.Join(os.Args[1:], " ")

		fmt.Println(SpOnGeBoB(input))
	}
}

// SpOnGeBoB formats a string in a very mocking way
func SpOnGeBoB(input string) (output string) {
	for i, v := range input {
		s := string(v)
		if i%2 == 0 {
			output += strings.ToLower(s)
		} else {
			output += strings.ToUpper(s)
		}
	}

	return
}
