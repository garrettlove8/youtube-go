package main

import "fmt"

/*
	- Go takes care of formatting, not us
	- Use the gofmt command
	- The -w flag will overwrite what's in you source code
*/

type person struct {
	name int
	age  int
}

func format() {
	newVar := "New variable"
	fmt.Println(newVar)
}
