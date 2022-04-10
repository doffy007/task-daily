package main

import "fmt"

func main() {
	var total = 100
	for n := 0; n < total; n++ {
		Keyword(int32(n))
	}
}

func Keyword(n int32) {
	if n%3 == 0 && n%2 == 1{
		fmt.Println(n, "- I Love Coding")
	} else if n%2 == 0 {
		fmt.Println(n, "- Berkualitas")
	} else if n%2 == 1{
		fmt.Println(n, "- Santai")
	}
}