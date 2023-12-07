package main

import (
	"fmt"

	_ "github.com/mypkg1/pkg1"
	_ "github.com/mypkg1/pkg1/pkg2"
)

func main() {
	// define int's a&b
	var a int = 10
	var b int = 20
	a1, b2 := 10, 20
	// define floats
	var c float64 = 67.54
	c1 := 3.14
	// define bool
	var a2 bool = true
	a3 := false
	a3 = true
	// define string, byte(8), rune(16)
	var hello string = "v1"
	hello = "v3"
	hello1 := "golang go"
	hello2 := `ŋ`
	for h := range hello2 {
		fmt.Printf("%T,%v\n", hello2[h], hello2[h])
	}
	// define a constant
	const const1 = 20
	const (
		v1 = 2
		v2 = "dgdgdg"
	)
	// const1 = 5 -- not possible

	// define a array 0,1,2,3,4,5
	var arr = [6]int{1,2} // length
	arr[2],arr[3]=3,4
	fmt.Println(arr[3]) // 0
	// define a slice
	sliceZeros := arr[3:] // length, capacity
	sliceZeros = append(sliceZeros, arr[:]...) // 0,0,0   1,2,3,4,0,0

	// cap = 4*32bits--> [32][32][32][32]  [32][32][32][32] 

	fmt.Println(arr[0],arr[2], len(sliceZeros), cap(sliceZeros))
	// define a struct & create it object.
	// define a map
	// define a Interface
	// define a channel
	fmt.Println(a, b, a1, b2, c, c1, a2, a3, hello1, hello, hello2, sliceZeros)
}
