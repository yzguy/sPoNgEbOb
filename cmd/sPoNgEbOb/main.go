package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/yzguy/sPoNgEbOb"
)

func main() {
	if len(os.Args) > 1 {
		input := strings.Join(os.Args[1:], " ")

		fmt.Println(sPoNgEbOb.Mock(input))
	}
}
