package main

import "fmt"

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
func findIndices(arr []int, target int) []int {

	var indexes []int
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if target == arr[i]+arr[j] && !contains(indexes, i) && !contains(indexes, j) {
				indexes = append(indexes, i)
				indexes = append(indexes, j)
			}
		}
	}

	return indexes

}

func main() {

	num := []int{2, 7, 9, 11, 15}

	target := 11

	fmt.Println(findIndices(num, target))

}
