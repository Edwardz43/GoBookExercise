package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	rotate(s)
	fmt.Println(s[:])
}

func rotate(s []int) {
	f := s[0]
	copy(s, s[1:])
	s[len(s)-1] = f
}
