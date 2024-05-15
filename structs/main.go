package main

import (
	"time"
)

type Employee struct {
	Name   string
	Number int
	Boss   *Employee
	Hired  time.Time
}

func main() {

	/*fmt.Println("Hello From Structs")
	var a Employee
	var e = Employee{
		Name:   "Piyush",
		Number: 3,
		Hired:  time.Now(),
	}
	e := Employee{
		Name:   "Piyush",
		Number: 3,
		Hired:  time.Now(),
	}

	b := &Employee{"Ayush", 1, nil, time.Now()}
	e.Boss = b

	a.Name = "Ayush"
	a.Number = 1
	a.Hired = time.Now()

	fmt.Printf("%T %+[1]v\n", a)
	fmt.Printf("%T %+[1]v\n", e)
	fmt.Printf("%T %+[1]v\n", b)
	c := map[string]*Employee{}
	c["Ayush"] = &Employee{
		Name:   "Ayush",
		Number: 2,
		Boss:   nil,
		Hired:  time.Now(),
	}
	c["Ayush"].Number++
	c["Piyush"] = &Employee{"Piyush", 1, c["Ayush"], time.Now()}
	fmt.Printf("%T %+[1]v\n", c["Ayush"])
	fmt.Printf("%T %+[1]v\n", c["Piyush"])*/
}
