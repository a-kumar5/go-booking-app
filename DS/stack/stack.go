package main

import "fmt"

type StringStack struct {
	data []string
}

func (s *StringStack) Push(x string) {
	s.data = append(s.data, x)
}

func (s *StringStack) Pop() string {
	if l := len(s.data); l > 0 {
		t := s.data[l-1]      // top
		s.data = s.data[:l-1] // head
		return t
	}
	panic("pop from empty stack")
}

func main() {
	fmt.Println("Hello From Stack")
	s := StringStack{}
	s.Push("Ayush")
	s.Push("Aish")
	s.Push("Ayu")
	t := s.Pop()
	fmt.Println(s)
	fmt.Println(t)
}
