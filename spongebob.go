package sPoNgEbOb

import "strings"

// Mock formats a string in a very mocking way
func Mock(input string) (output string) {
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
