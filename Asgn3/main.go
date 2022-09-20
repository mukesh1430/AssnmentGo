package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getBinary(n int) []int {

	arr := make([]int, 0, 0)

	for i := 0; i <= n; i++ {

		b := int64(i)

		binaryRep := strconv.FormatInt(b, 2)

		arr = append(arr, strings.Count(binaryRep, "1"))

	}

	return arr

}

func main() {

	finalArr := getBinary(5)

	fmt.Println(finalArr) //[0 1 1 2]

}
