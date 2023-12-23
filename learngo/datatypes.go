package learngo

import "fmt"

func dt1() {
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
	var arr = [6]int{1, 2} // length
	arr[2], arr[3] = 3, 4
	fmt.Println(arr[3]) // 0
	// define a slice
	sliceZeros := arr[3:]                      // length, capacity
	sliceZeros = append(sliceZeros, arr[:]...) // 0,0,0   1,2,3,4,0,0
	// arr --> 1234 --> append(sliceZeros, 1,2,3,4)
	// cap = 4*32bits--> [32][32][32][32]  [32][32][32][32]

	fmt.Println(arr[0], arr[2], len(sliceZeros), cap(sliceZeros))
	// define a struct & create it object.

	type abcd struct {
		hello  string
		hello1 int
		hello2 bool
		s1     []int
		s2     interface{} //empty interface
	}

	obj := abcd{
		hello:  "hello str",
		hello1: 2,
		s2:     "string"}
	obj2 := []abcd{
		{"h1", 2, true, []int{1, 2, 3}, 5},
		{"h2", 2, true, []int{1, 2, 3}, 5},
	}
	fmt.Println(obj)
	obj2 = append(obj2, abcd{"h3", 2, true, []int{1, 2, 3}, 5}, obj)
	fmt.Println(obj2)

	// define a map
	var menu map[string]int
	menu = map[string]int{
		"noodels": 50,
		"dosa":    60,
	}
	menu["dosa"] = 70
	menu["idli"] = 20
	fmt.Println(menu)
	// define a Interface
	var int1 interface{} //empty interface
	//
	int1 = 1
	int1 = "abcd"
	fmt.Println(int1)

	fmt.Println(a, b, a1, b2, c, c1, a2, a3, hello1, hello, hello2, sliceZeros)

	//interfaces
	var stud1 = stud{1, "pawan", "10th", 50}
	marks := stud1.Student("pawan", 1)
	marks2 := stud1.Student2()
	fmt.Println(marks, marks2)
	printer(stud1) // interface calling
	// define a channel
	// ->[int]->

	ch1 := make(chan int, 1) // buffered
	ch1 <- 5                 // [5]
	data1 := <-ch1           //5<-[]
	ch1 <- 7                 // [7]
	data2 := <-ch1           // 7<-[]
	close(ch1)
	data3 := <-ch1        // 0
	ch2 := make(chan int) // unbuffered channel
	close(ch2)
	// ch2 <- 10 | "<-ch2"
	fmt.Println(data1)
	fmt.Println(data2)
	fmt.Println(data3)

}

type School interface { // aim School --> student, teacher, xyz
	Student(string, int) int //
	ClassName(int) string    //id --> 10th
}

type stud struct {
	id        int
	name      string
	className string
	marks     int
}

func (s stud) Student(s_name string, id int) int { // s is known as reciver for the function
	fmt.Println(s_name, id)
	return 20
}

func (s stud) Student2() int { // you can call structobject.Student2()
	fmt.Println(s.name, s.id)
	return s.marks
}

func (s stud) ClassName(id int) string {
	fmt.Println(id)
	return "10th"
}

// /
func printer(s School) { // 2nd use of interfaces, we can pass stud object in place of School
	marks := s.Student("general_user", 123)
	fmt.Println(marks)
}
