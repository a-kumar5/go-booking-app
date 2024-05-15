package main

import (
	"fmt"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	count_s := map[string]int{}
	count_t := map[string]int{}

	for _, c := range s {
		if _, count := count_s[string(c)]; count == false {
			count_s[string(c)] = 0
		}
		count_s[string(c)] += 1
	}
	for _, c := range t {
		if _, count := count_s[string(c)]; count == false {
			count_s[string(c)] = 0
		}
		count_t[string(c)] += 1
	}

	for key, _ := range count_s {
		_, ok := count_t[key]

		if ok == false || count_t[key] != count_s[key] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Hello, 世界")
	var s = "anagram"
	var t = "nagaram"
	fmt.Println(isAnagram(s, t))
}
