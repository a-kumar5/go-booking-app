package main

import (
	"fmt"
	"math"
)

type Point struct{ X, Y float64 }
type Path []Point
type IntList struct {
	value int
	Tail  *IntList
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

func main() {
	a := IntList{1, &IntList{3, nil}}
	b := IntList{2, &a}
	fmt.Println((&b).Sum())
	fmt.Println(b.Sum())
	/*
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
