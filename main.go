package main

import (
	"errors"
	"fmt"

	"github.com/mypkg1/learngo"
)

func divide(a, b, index int) (result int) {
	defer func() {
		if r := recover(); r != nil { // if panic occurs recover will return r that is not nil, and we are printing it.
			fmt.Println("There was a panic due to : ", r)
			// result = -1
			// os.Exit(1)
		}
	}()
	if b == 0 {
		result = -1 // hardcoding the results
		panic("Hey, please don't send b=0..!")
	}
	result = a / b
	arr1 := []int{1, 2, 3, 4}
	if index > 3 {
		panic("Hey, please don't send index>3..!")
	}
	fmt.Println(arr1[index])
	return
}

func handelErrorsWithPanic() {
	fmt.Println("Division output : ", divide(4, 2, 7))
	fmt.Println("Division output : ", divide(1, 0, 0))
	fmt.Println("Division output : ", divide(2, 2, 100))
}

func divideV2(a, b, index int) (result int, err error) {
	if b == 0 {
		return -1, errors.New("hey, please don't send b=0")
	}
	result = a / b //4/2 = 2

	arr1 := []int{1, 2, 3, 4}
	if index > 3 {
		return -1, errors.New("hey, please don't send index>3")
	}
	fmt.Println(arr1[index])

	return result, nil
}

func main() {
	// standard way of handling errors
	logger := learngo.GetSingleInstanceLogger()
	resultoutput, err := divideV2(4, 2, 7) // -1,"hey, please don't send index>3"
	if err == nil {
		fmt.Println("Division output : ", resultoutput)
		logger.Info("Executed without errors.")
	} else {
		fmt.Printf("there was an error : %v output : %v \n", err.Error(), resultoutput)
		logger.Error("Executed with ERRORS ", "error_value", err.Error())
	}
	logger.Info("Done with code.")
}
