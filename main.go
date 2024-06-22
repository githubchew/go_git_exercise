package main

import (
	"fmt"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	var todo []string
	todo = append(todo, "sing")
	todo = append(todo, "cook")
	todo = append(todo, "sow")
	todo = append(todo, "plank")
	todo = append(todo, "go","eat","play")
	tomorrow := []string{ "see mom","learn golang"}

	todo = append(todo, tomorrow...)
	s.Show("todo", todo)

}
