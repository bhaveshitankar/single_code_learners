package learngo

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func printExitMsg() {
	// time.Sleep(4 * time.Second)
	fmt.Println("I am exiting...")
}

func goroutine1() {
	defer wg.Done()
	defer printExitMsg()
	// wg.Done() // removing 1-wait from stack // w=1, w=0
	fmt.Println("hello from first go-routine")
	// panic/ returned
	if true {
		return
	}
	defer printExitMsg() // not reachable defer
	// wg.Done()
}

// using wait groups
func Concurrent() {
	// runtime.GOMAXPROCS(3) // parallel
	wg.Add(2) // adding 2-waits to stack // 3-w

	go goroutine1() // not getting enough time to complete | --> new thread
	go goroutine1() // not getting enough time to complete | --> new thread

	// giving 1 sec delay before main exits
	// time.Sleep(1 * time.Second)
	// wg.Done() // 2-w
	wg.Wait() // waiting till waits become zero // is w == 0
}

func Parallel() {
	runtime.GOMAXPROCS(3) // parallel
	wg.Add(3)             // adding 2-waits to stack // 3-w

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
