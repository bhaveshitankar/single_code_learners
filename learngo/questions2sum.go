package learngo

import "fmt"

func solving2sum(arr []int, target int) {
	// W.A.P , you are give an array arr = []int{1,2,3,4}, target=5, find the indexes in array which sums up to target
	// for example arr[0]+arr[3] = 5 --> o/p [0,3], [1,2]
	m1 := map[int]int{}
	for idx := range arr {
		m1[arr[idx]] = idx // m1[1]=0
	}
	for idx := range arr {
		valuetobefound := target - arr[idx]
		if v, ok := m1[valuetobefound]; ok {
			fmt.Println("fist-index: ", idx, "second-index: ", v)
			break
		}
	}
}
