package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func multiply(a, b int) int {
	return a * b
}

func main() {
	fmt.Println(multiply(1, 2))
	test, test1 := lenAndUpper("TEST")
	fmt.Println(test, test1)
}
