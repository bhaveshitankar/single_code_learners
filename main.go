package main

import "fmt"

func chopping() { // t=0
	fmt.Println("I am chopping now...") // t=1 started printing
	ch <- 1                             // t=2
} // t=3

func cooking() { // t=0
	fmt.Println("I am cooking now...") // t=1 started printing.
	ch <- 2                            //t=2
} //t=3

var ch = make(chan int, 2)

func main() {
	defer close(ch)
	go chopping()          // t=0s calling chopping
	go cooking()           // t=0s calling cooking
	dataOfchopping := <-ch // t=2
	dataOfcooking := <-ch  //t=2
	fmt.Println(dataOfchopping, dataOfcooking)
} // t=2
