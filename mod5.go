package main

import "fmt"

func main() {
	var a = "goodafernoon"

	fmt.Print("hello everybody")

	fmt.Print(a)

	fmt.Print("hello everybody", a)
	fmt.Print("hello everybody", a, "and have a nice day")
	fmt.Print("and have a nice day")

}
