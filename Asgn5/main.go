package main

import (
	"fmt"
	"time"
)

func routine1(c chan<- int) {

	for i := 1; i < 10; i++ {
		c <- i
	}

}

func routine2(d chan int) {

	for {

		num := <-d

		if num == 1 {
			fmt.Printf("1 is neither prime nor composite number.\n")
		} else {
			var isPrime bool = true
			for i := 2; i < num; i++ {

				if num%i == 0 {
					isPrime = false
					break
				}
			}
			if isPrime {

				fmt.Printf("%d is prime number \n", num)

			} else {
				fmt.Printf("%d is not prime number \n", num)

			}
		}

	}

}

func main() {

	d := make(chan int)
	go routine1(d)
	go routine2(d)

	time.Sleep(time.Second)

	// fmt.Println("")

}
