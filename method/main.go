package main

import (
	"fmt"
	"image/color"
	"io"
	"math"
	"os"
)

type Point struct{ X, Y float64 }
type Path []Point
type IntList struct {
	value int
	Tail  *IntList
}

// Values maps a string key to a list of values.
type Values map[string][]string

type ColoredPoint struct {
	Point
	Color color.RGBA
}

// Get returns the first value associated with the given key,
// or "" if there are none.
func (v Values) Get(key string) string {
	if vs := v[key]; len(vs) > 0 {
		return vs[0]
	}
	return ""
}
func (v Values) Add(key, value string) {
	v[key] = append(v[key], value)
}

func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	return list.value + list.Tail.Sum()
}

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

// Methods with a Pointer Receiver
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

type ByteCounter int

func (b *ByteCounter) Write(p []byte) (int, error) {
	l := len(p)
	*b += ByteCounter(l)
	return l, nil
}

func main() {
	var c ByteCounter
	f1, _ := os.Open("a.txt")
	//f2, _ := os.Create("out.txt")
	f2 := &c

	n, _ := io.Copy(f2, f1)
	fmt.Println("copied", n, "bytes")
	fmt.Println(c)
	/*
		var cp ColoredPoint
		cp.X = 1
		fmt.Println(cp.Point.X) // 1
		red := color.RGBA{255, 0, 0, 255}
		blue := color.RGBA{0, 0, 255, 255}
		var p = ColoredPoint{Point{1, 1}, red}
		var q = ColoredPoint{Point{5, 4}, blue}
		fmt.Println(p.Distance(q.Point)) // 5
		p.ScaleBy(2)
		q.ScaleBy(2)
		fmt.Println(p.Distance(q.Point)) // 10
		//p.Distance(q) // compile error - ./main.go:79:13: cannot use q (variable of type ColoredPoint) as Point value in argument to p.Distanc

			m := url.Values{"lang": {"en"}}
			fmt.Println(m.Get("lang"))
			m.Add("item", "1")
			m.Add("item", "2")
			fmt.Println(m.Get("q"))
			fmt.Println(m.Get("item"))
			fmt.Println(m["item"]

				a := IntList{1, &IntList{3, nil}}
				b := IntList{2, &a}
				fmt.Println((&b).Sum())
				fmt.Println(b.Sum())
					p := Point{1, 2}
					q := Point{4, 6}
					fmt.Println(Distance(p, q))
					fmt.Println(p.Distance(q))
					perim := Path{
						{1, 1},
						{5, 1},
						{5, 4},
						{1, 1},
					}
					fmt.Println(perim.Distance())

					r := &Point{1, 2}
					r.ScaleBy(2)
					fmt.Println(*r)
					s := Point{1, 2}
					sptr := &s
					sptr.ScaleBy(2)
					fmt.Println(s)
	*/
}
