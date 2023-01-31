package main

import (
	"fmt"
	"io"
)

func main() {
	var ans int = 0
	for {
		var str int
		_, err := fmt.Scanf("%s\n", &str)
		if err != io.EOF {
			ans++
		} else {
			break
		}
	}
	fmt.Println(ans)
}
