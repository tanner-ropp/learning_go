package main

import "fmt"

func main() {
	// this const is untyped integer, just like literals are untyped in go
	const value = 42
	// because "value" is an integer literal and defaults to type int, we can use the shorthand and i is inferred to be of type int
	i := value
	// type has to specified here since it would infer int by default
	var f float64 = value
	fmt.Println(i, f)
}
