package main

import (
	"fmt"
	"two/internal/names"
)

func main() {
	// format()
	// Comments()
	person := names.New("Glove")
	fmt.Println("Person's name: ", person.Name())
	return
}
