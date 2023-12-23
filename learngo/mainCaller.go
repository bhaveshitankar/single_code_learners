package learngo

import "fmt"

func callall() {
	// conditionalOperators()
	// callerForAllFunctions()
	// write a func to check if a key exists in a map.
	m1 := map[string]int{"dosa": 55, "idli": 32}
	exists := checkMap("dosa", m1)
	exists2 := checkMap("noodels", m1)
	fmt.Println(exists, exists2)
	// W.A.P to, update dosa, noodels prices if exist in m1 "menuUpdator"
	// W.A.P to, count number of times each character is repeated. ["a":4,"b":1...]
	s1 := "abcdasaakkefg"
	m2 := map[string]int{}
	/*for _, v := range s1 {
		if checkMap(string(v), m2) {
			// character exists
			// m2[string(v)] = m2[string(v)] + 1
			m2[string(v)] += 1
		} else {
			m2[string(v)] = 1
		}
	}*/
	for _, v := range s1 {
		if _, ok := m2[string(v)]; ok { // _, ok := m2[0]
			// character exists
			// m2[string(v)] = m2[string(v)] + 1
			m2[string(v)] += 1
		} else {
			m2[string(v)] = 1
		}
	}
	fmt.Println(m2)
	// W.A.P , you are give an array arr = []int{1,2,3,4}, target=5, find the indexes in array which sums up to target
	// for example arr[0]+arr[3] = 5 --> o/p [0,3], [1,2]
	arr := []int{1, 2, 3, 4}
	target := 5
	// brutforce approch. n*n = O(n^2) --> n --> nlog(n)
	for idx1 := range arr {
		checkBreak := false
		for j := idx1 + 1; j < len(arr); j++ {
			if target == arr[idx1]+arr[j] {
				fmt.Println("fist-index: ", idx1, "second-index: ", j)
				checkBreak = true
				break
			}
		}
		if checkBreak {
			break
		}
	}
	// optimized solution:
	fmt.Println("optimized solution output : ")
	solving2sum(arr, target)
}
