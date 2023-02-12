package main

import "fmt"

func main() {
	n := 3
	//j := &n
	testPointer(&n)
	fmt.Println(n)

	k := testTwoPointer(n)
	fmt.Println(*k)
}

func testPointer(n *int) {
	*n = *n * *n
}

func testTwoPointer(n int) *int {
	v := n * n
	return &v
}
