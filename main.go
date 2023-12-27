package main

import "fmt"

func goroutine1() {
	fmt.Println("hello from first go-routine")
}

func main() {
	go goroutine1()
	go goroutine1()
}
