package main

import "fmt"

func main() {
	var a map[string]interface{}
	b := map[string]interface{}{}

	fmt.Println(a == nil, len(a) <= 1)

	fmt.Println(b == nil, len(b) <= 1)
}
