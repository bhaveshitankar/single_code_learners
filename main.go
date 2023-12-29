package main

import "fmt"

func goroutine1() {
	fmt.Println("hello from first go-routine")
}

func main() {
	go goroutine1()
	go goroutine1()
}

// Schedulers :- Plan your execution, decides what to run and when to run, where to run on cpu
// Paralyzation :- 2 ppl --> chopping, cooking --> parallel task.
// Concurrency :- 1 ppl --> chops, mix-cooks, chops, mix-cooks --> making sure both tasks are progressing.
// runtime.MAXPROCKS = 8
// by telling Schedulers to use mu;ti core we convert Concurrency to Paralyzation
// if main exits all go routine will exit in what-ever state they are.
