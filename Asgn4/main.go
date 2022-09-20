package main

import (
	"fmt"
	"time"
)

var alphaArr = []string{"A", "B", "C", "D", "E"}
var intArr = []int{1, 2, 3, 4, 5}

func routine1(c chan<- string) {

	for i := 0; i < len(alphaArr); i++ {
		c <- alphaArr[i]
	}

}

func routine2(d chan<- int) {

	for i := 0; i < len(intArr); i++ {
		d <- intArr[i]
	}

}

func PrintFS(c <-chan string, d <-chan int) {
	for {
		charOt := <-c
		intOt := <-d

		fmt.Printf("%d%s", intOt, charOt)
	}

}

func main() {
	var c chan string = make(chan string)
	d := make(chan int)
	go routine1(c)
	go routine2(d)

	go PrintFS(c, d)

	time.Sleep(time.Second)

	fmt.Println("")

}
