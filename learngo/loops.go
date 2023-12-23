package learngo

import "fmt"

func loops() {
	fmt.Println("Normal for loop")
	for i := 0; i < 10; i++ {
		fmt.Println(i * i)
	}

	fmt.Println("Normal while loop")
	i := 0
	for {
		if i == 10 { // 1 ==10
			break // line - 160
		}
		if i%2 == 0 { // 1%2 ->1
			i++
			continue // line 150 , ie. next iteration.
		}
		fmt.Println(i * i)
		i++
	} // 1,9,25,49,...,81

	fmt.Println("Normal array loop")
	arr1 := [5]int{1, 2, 34, 5, 6}
	for i := 0; i < len(arr1); i++ {
		fmt.Println(i, arr1[i])
	}

	fmt.Println("array loop with range")
	for idx, value := range arr1 {
		fmt.Println(idx, value)
	}

	fmt.Println("map loop with range")
	map1 := map[string]string{"key1": "value1", "key2": "VALUE2"}
	for key, value := range map1 {
		fmt.Println(key, value)
	}

	for _, value := range map1 {
		fmt.Println(value)
	}

	for idx := range arr1 {
		fmt.Println("index: ", idx, arr1[idx])
	}

	for k := range map1 {
		fmt.Println("key: ", k, map1[k])
	}
}
