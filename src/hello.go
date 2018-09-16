package main

import "fmt"
import "strings"

func hello() string {
	fmt.Println(strings.ToUpper("calling bar"))

	foo := func() string {
		return "foo"
	}

	return foo()
}
