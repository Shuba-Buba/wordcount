package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	phrase := os.Args[1]
	words := strings.Fields(phrase)
	fmt.Println(len(words))
}
