package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
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
		//cat(file)
		//fmt.Println("\n----------------------------------")
		//size(file)
		//fmt.Println("\n----------------------------------")
		wordCount(file)
	}
}

func cat(file *os.File) {
	if _, err := io.Copy(os.Stdout, file); err != nil {
		fmt.Fprint(os.Stderr, err)
		return
	}
	return
}
func size(file *os.File) {
	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return
	}
	fmt.Println("The file has", len(data), "bytes")
	return
}

func wordCount(file *os.File) {
	var lc, wc, cc int
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		s := scan.Text()
		wc += len(strings.Fields(s))
		cc += len(s)
		lc++
	}
	fmt.Printf(" %7d %7d %7d \n", lc, wc, cc)
}
