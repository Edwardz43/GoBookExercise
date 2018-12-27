package main

import "fmt"

func main() {
	s := []string{"1", "2", "2", "3", "3", "3", "4"}

	s = unique(s)

	fmt.Println(s)
}

func unique(s []string) []string {

	var res []string

	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			continue
		}
		res = append(res, s[i])
	}
	res = append(res, s[len(s)-1])
	return res
}
