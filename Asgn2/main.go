package main

import "fmt"

func add(a int, b int) int {

	if b == 0 {
		return a
	}

	for i := 1; i <= b; i++ {
		a++

	}

	return a

}

func main() {

	a := 5
	b := 9

	sum := add(a, b)

	fmt.Println(sum)

}
