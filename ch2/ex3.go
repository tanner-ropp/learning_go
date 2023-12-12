package main

import "fmt"

func main() {
	// since we are making multiple variables we can use a declaration list
	var (
		b      byte   = 255
		smallI int32  = 2147483647
		bigI   uint64 = 18446744073709551615
	)

	b = b + 1
	smallI = smallI + 1
	bigI = bigI + 1

	// because we added 1 to the max values of each of these types, we should see them wrap back around to their smallest values
	fmt.Println(b, smallI, bigI)
}
