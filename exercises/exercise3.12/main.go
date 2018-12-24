package main

import (
	"fmt"
	"os"
)

func main() {
	//fmt.Println(check("asdsadasd", "dasdasdsa"))

	if a := os.Args[1:]; len(a) <= 0 {
		fmt.Printf("Too few arguments !\n")
		os.Exit(1)
	}

	if a := os.Args[1:]; len(a) > 2 {
		fmt.Printf("Too many arguments !\n")
		os.Exit(1)
	}

	fmt.Println(check(os.Args[1], os.Args[2]))
}

func check(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	t1 := make(map[rune]int)

	for _, c := range s1 {
		t1[c]++
	}

	t2 := make(map[rune]int)
	for _, c := range s2 {
		t2[c]++
	}

	for k, v := range t1 {
		if t2[k] != v {
			return false
		}
	}

	for k, v := range t1 {
		if t1[k] != v {
			return false
		}
	}

	return true
}
