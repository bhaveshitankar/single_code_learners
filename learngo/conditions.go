package learngo

import "fmt"

func conditionalOperators() {
	// if , switch case,

	// switch inp := 2;inp {
	inp := 1
	switch inp {
	case 1:
		fmt.Println("Thanks for choosing English..!")
		fmt.Println("To change press 9")
	case 2:
		fmt.Println("Thanks for choosing Hindi..!")
		fmt.Println("To change press 8")
	default:
		fmt.Println("In valid input")
	}

	inp = 2
	if inp == 1 {
		fmt.Println("Thanks for choosing English..!")
		fmt.Println("To change press 9")
	} else if inp == 2 {
		fmt.Println("Thanks for choosing Hindi..!")
		fmt.Println("To change press 8")
	} else {
		fmt.Println("In valid input")
	}
}
