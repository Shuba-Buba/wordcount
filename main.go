package main

import (
	"fmt"
	"io"
)

func main() {
	var ans int = 0
	for {
		var str string
		_, err := fmt.Scanf("%s", &str)
		if err != io.EOF {
			ans++
		} else {
			break
		}
	}
	fmt.Println(ans)
}
