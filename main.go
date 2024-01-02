package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func goroutine1() {
	fmt.Println("hello from first go-routine")
	wg.Done() // removing 1-wait from stack // w=1, w=0
}

func main() {
	// runtime.GOMAXPROCS(3) // parallel
	wg.Add(3) // adding 2-waits to stack // 3-w

	go goroutine1() // not getting enough time to complete | --> new thread
	go goroutine1() // not getting enough time to complete | --> new thread

	// giving 1 sec delay before main exits
	// time.Sleep(1 * time.Second)
	wg.Done() // 2-w
	wg.Wait() // waiting till waits become zero // is w == 0
}

// Schedulers :- Plan your execution, decides what to run and when to run, where to run on cpu
// Paralyzation :- 2 ppl --> chopping, cooking --> parallel task.
// Concurrency :- 1 ppl --> chops, mix-cooks, chops, mix-cooks --> making sure both tasks are progressing.
// runtime.MAXPROCKS = 3
// by telling Schedulers to use multi core we convert Concurrency to Paralyzation
// if main exits all go routine will exit in what-ever state they are.
// Paralyzation : used for cpu intense calculations.
// Concurrency : used for external input influenced programming.
