package main

import (
	"errors"
	"fmt"
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

func divideV2(a, b, index int) (result int, err error) {
	if b == 0 {
		result, err = 0, errors.New("trying to devide by zero...")
		return
	}
	result = a / b
	arr1 := []int{1, 2, 3, 4}
	fmt.Println(arr1[index])
	return
}

func main() {
	fmt.Println("Division output : ", divide(4, 2, 7))
	fmt.Println("Division output : ", divide(1, 0, 0))
	fmt.Println("Division output : ", divide(2, 2, 100))
}
