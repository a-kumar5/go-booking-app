package main

import (
	"fmt"
	"io"
	"os"
)

//"fmt"
//"hello"
//"os"

func main() {
	//fmt.Println(hello.Say(os.Args[1:]))
	//var w = [...]int{1, 2, 3}
	//var x = []int{0, 0, 0}
	//y := hello.Do(w, x)
	//fmt.Println(w, x, y)
	// sorting - map examples

	// hello.Sort()
	for _, fname := range os.Args[1:] {
		file, err := os.Open(fname)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		if _, err := io.Copy(os.Stdout, file); err != nil {
			fmt.Fprint(os.Stderr, err)
			continue
		}
		file.Close()
	}
}
