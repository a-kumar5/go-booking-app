package hello

import (
	"strings"
)

func Say(names []string) string {
	if len(names) == 0 {
		names = []string{"world"}
	}
	return "Hello, " + strings.Join(names, ", ") + "!"
}

func Do(a [3]int, b []int) []int {
	// a = b //syntax error
	a[0] = 4
	b[0] = 3

	c := make([]int, 5)
	c[4] = 42
	copy(c, b)
	return c
}
