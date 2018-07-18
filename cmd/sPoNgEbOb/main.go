package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/yzguy/sPoNgEbOb"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(sPoNgEbOb.Mock(scanner.Text()))
	}
}
