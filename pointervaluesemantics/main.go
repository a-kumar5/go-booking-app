package main

import "fmt"

type user struct {
	name  string
	count int
}

func addTo(u *user) {
	u.count++
}

func main() {
	fmt.Println("Hello From pointervaluesemantics")
	users := []user{{"alice", 0}, {"bob", 0}}

	alice := &users[0]

	amy := user{"amy", 0}

	users = append(users, amy)

	addTo(alice)
	fmt.Printf("%p\n", alice)
	fmt.Printf("%p\n", &alice)
	fmt.Println(*alice)
	fmt.Println(users)
	fmt.Printf("%p\n", users)
}
