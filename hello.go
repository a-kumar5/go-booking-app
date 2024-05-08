package hello

import (
	"strings"
	"bufio"
	"os"
	"fmt"
	"sort"
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

func Sort(){
	scan := bufio.NewScanner(os.Stdin)
	words := make(map[string]int)
	scan.Split(bufio.ScanWords)
	for scan.Scan(){
		words[scan.Text()]++
	}
	fmt.Println(len(words), "unique words")

	type kv struct {
		key string
		val int
	}

	var ss []kv

	for k, v := range words {
		ss = append(ss, kv{k, v})
	}

	fmt.Println(ss)

	sort.Slice(ss, func(i, j int) bool{
		return ss[i].val > ss[j].val
	})
	fmt.Println(ss)

	for _, s := range ss {
		fmt.Println(s.key, "appears", s.val, "times")
	}

}
