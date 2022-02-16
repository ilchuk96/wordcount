package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		return
	}
	line := args[1]
	worlds := strings.Fields(line)
	fmt.Println(len(worlds))
}
