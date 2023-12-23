package learngo

import "fmt"

func f1() {
	i := 1
	if i == 1 {
		return
	} else {
		fmt.Println("Hi from f1")
	}
	return
}

func f1WithParams(i int) {
	// i := 1
	if i == 1 {
		return
	} else {
		fmt.Println("Hi from f1WithParams")
	}
	return
}

func f1WithParamsAndReturn(i int, j int) (int, int, int) {
	return i, j, i * j
}

func f1WithParamsAndReturnShortcut(i, j int, str1, str2 string) (int, int, int) {
	fmt.Println(str1)
	fmt.Println(str2)
	return i, j, i * j
}

func callerForAllFunctions() {

	fmt.Println("\nrunning f1")
	f1()

	fmt.Println("\nrunning f1WithParams")
	f1WithParams(2)

	fmt.Println("\nrunning f1WithParamsAndReturn")
	i, j, mult := f1WithParamsAndReturn(2, 3)
	fmt.Println(i, j, mult)

	fmt.Println("\nrunning f1WithParamsAndReturnShortcut")
	i, j, mult = f1WithParamsAndReturnShortcut(4, 5, "hello World", "func of f1WithParamsAndReturnShortcut")
	fmt.Println(i, j, mult)
}
