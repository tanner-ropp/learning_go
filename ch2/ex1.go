package main

import "fmt"

func main() {
	// int type is inferred
	var i = 20
	// conversion has to be done manually, then the new type is inferred by assignment
	var f = float64(i)
	fmt.Println(i, f)
}
