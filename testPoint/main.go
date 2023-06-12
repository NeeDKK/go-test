package main

import (
	"fmt"
	"strings"
)

func main() {
	ss := []string{"internal", "public", "aaa", "ccc"}
	for i := 0; i < len(ss); i++ {
		if strings.Contains(ss[i], "internal") {
			ss = append(ss[:i], ss[i+1:]...)
			i--
		}
	}
	fmt.Print(ss)
}
