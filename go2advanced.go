/*
These are comments in go
*/

go mod init example.com/dev  // creats a go.mod file in wdir
go mod tidy // updates mod file.

go get //

// GOPATH Not needed

## Variables / constants:
    var variablename type = value
    variablename := value   /// dosent work for global var

    var student1 string = "John" //type is string
    var b1 bool = true
    var student2 = "Jane" //type is inferred
    x := 2 //type is inferred
    var a, b = 6, "Hello"
    c, d := 7, "World!"
    a int // has value 0
    // int,uint, int8..int64
    // float32, float64 

    //constants
    const PI = 3.14
    const (
            A int = 1
            B = 3.14
            C = "Hi!"
        )

    // Arrays
    var array_name = [length]datatype{values} // here length is defined
    var array_name = [...]datatype{values} // here length is inferred
    var arr1 = [...]int{1,2,3,4}
    arr1 := [...]int{1,2,3,4} // len(arr1) = 4

        //slices  (same as arrays but with variable sizes)
        slice_name := []datatype{values}
        slice_name := make([]type, length, capacity)
        slice_name := myArr[start:end]

        myslice1 := make([]int, 5, 10)
        myslice2 := []int{10, 11, 12, 13, 14,15}
        arr1 := [6]int{10, 11, 12, 13, 14,15}
        slice_from_arr1 := arr1[2:4]
        
        //appending in slice , append(slice_name,el1,el2,...),len(slice_name),cap(slice_name),copy(dest, src)
        slice_name = append(slice_name, element1, element2, ...)
        slice3 = append(slice1, slice2...)
        myslice2 = append(myslice2,16)

        
    //Structs
    type struct_name struct {
        member1 datatype;     //; is optional
        member2 datatype;
        member3 datatype;
        ...
    }
    //eg
    type myStruct struct{
        var1 string
        var2 int
    }
    //initilization 
    ms1 := myStruct{"xyz"} // ms1 = {"xyz",0}
    ms2 := myStruct{"xyz",8} // ms1 = {"xyz",8}
    var ms3 myStruct
    ms3.var1="xyz";ms3.var2=9

    var pa *Student    // pa == nil
    pa = new(Student)  // pa == &Student{"", 0}
    pa.Name = "Alice"  // pa == &Student{"Alice", 0}

    pb := &Student{    // pb == &Student{"Bob", 8}
            Name: "Bob",
            Age:  8,
        }


    //maps
    var a = map[KeyType]ValueType{key1:value1, key2:value2,...}
    b := map[KeyType]ValueType{key1:value1, key2:value2,...}

    var a = make(map[KeyType]ValueType)
    b := make(map[KeyType]ValueType)

    var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
    b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}
    var a = make(map[string]string)

        //looping in array 
        for idx, value := range fruits
        //loop in slices
        for idx, value := range myslice2
        //looping in map
        for key, element := range employee

    //Interfaces--> they are general purpose methods sets which are implemented by different types.
    var manyType interface{}
    type Employee interface {
        PrintName() string                // Method with string return type
        PrintAddress(id int)              // Method with int parameter
        PrintSalary(b int, t int) float64 // Method with parameters and return type
    }
        //lets make 2 general purpose interfaces 
        type Polygons interface {
            Perimeter()
        }
        type Object interface {
            NumberOfSide()
        }
        // lets create some types that will implement above methods
        type pentagon int
        type square int
        // lets implement above interface methods for type pentagon type
        func (p pentagon) Perimeter() {
            fmt.Println(p*5)
        }
        func (p pentagon) NumberOfSide() {
            fmt.Println(5)
        }
        // lets use it
        var obj Object = Pentagon(50)
        obj.NumberOfSide()
        var pent Pentagon = obj.(Pentagon)
        pent.Perimeter()

    // channels
    Unbuffered := make(chan int) // Unbuffered channel of integer type
    buffered := make(chan int, 10)	// Buffered channel of integer type
    goroutine1 <- "Australia" // Send a string through the channel.
    data := <-goroutine1 // Receive a string from the channel.

### Slices capacity test
    
        package main
        import ("fmt")

        func main() {
            // see capas changing
            arr := []struct{}{}
            oldCap := 0
            for i := 0; i < 100; i++ {
                arr = append(arr, struct{}{})
                if cap(arr) != oldCap {
                    oldCap = cap(arr)
                    fmt.Println("arr", cap(arr))
                }
        }
        // just a switch case example
        day := 1
        switch day {
            case 1,5:
                fmt.Println("Monday")
            case 2:
                fmt.Println("Tuesday")
            }
        }

### typecasting
        strconv.ParseFloat(s string, bitSize int) (float64, error)
        strconv.Atoi(s string) (i int, err error)
        strconv.ParseInt(s string, base int, bitSize int) (i int64, err error)
        b1, _ := strconv.ParseBool(s1)
        s := fmt.Sprintf("%v", b)  // OR // strconv.FormatInt/FormatBool/FormatFloat-->(var s string = strconv.FormatFloat(f, 'E', -1, 32))
        int64(int32(int16(int(a))))
        reflect.TypeOf(s)

### Loops/ifs/functions

    // For statement with array
    fruits := [3]string{"apple", "orange", "banana"}

    for idx, _ := range fruits {
        fmt.Printf("%v\n", idx)
    }

    // writing assingments in if statement
    if file, err := os.Open(filename); err != nil {
     ....// file var is avilable here only, same works with for and switch
    }

    //function params and returns
    func myFunction(x int, y string) (result int, txt1 string) {
        result = x + x
        txt1 = y + " World!"
        return  //OR//return result,txt1//OR//return 55,"xyz"//order matters.
    }
    // OR
    func myFunction(x int, y string) (int,string) {
        result := x + x
        txt1 := y + " World!"
        return result, txt1
    }
    //vardict functions
    func variadicExample(s ...string){ } --> variadicExample("s1","s2")


## Output
    fmt.Print()
    fmt.Println()
    fmt.Printf()

    fmt.Printf("j has value: %v and type: %T", j, j) #%v is value, %T is type




// Panic and Recovery
package main
import (
	"errors"
	"fmt"
)
var result = 1
func chain(n int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	if n == 0 {
		panic(errors.New("Cannot multiply a number by zero"))
	} else {
		result *= n
		fmt.Println("Output: ", result)
	}
}
func main() {
	chain(5)
	chain(2)
	chain(0)
	chain(8)
}
/*
Output:  5
Output:  10
Cannot multiply a number by zero
Output:  80*/


// Accepting addresses in interfaces
package main
import "fmt"

type Book struct {
	author, title string
}
type Magazine struct {
	title string
	issue int
}

func (b *Book) Assign(n, t string) {
	b.author = n
	b.title = t
}
func (b *Book) Print() {
	fmt.Printf("Author: %s, Title: %s\n", b.author, b.title)
}

func (m *Magazine) Assign(t string, i int) {
	m.title = t
	m.issue = i
}
func (m *Magazine) Print() {
	fmt.Printf("Title: %s, Issue: %d\n", m.title, m.issue)
}

type Printer interface {
	Print()
}

func main() {
	var b Book                                 // Declare instance of Book
	var m Magazine                             // Declare instance of Magazine
	b.Assign("Jack Rabbit", "Book of Rabbits") // Assign values to b via method
	m.Assign("Rabbit Weekly", 26)              // Assign values to m via method

	var i Printer // Declare variable of interface type
	fmt.Println("Call interface")
	i = &b    // Method has pointer receiver, interface does not
	i.Print() // Show book values via the interface
	i = &m    // Magazine also satisfies shower interface
	i.Print() // Show magazine values via the interface
}

### Deffered function
package main
import "fmt"
func main() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}
// prints 4 3 2 1 0 

#### Goroutines with waitgroups

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

// WaitGroup is used to wait for the program to finish goroutines.
var wg sync.WaitGroup

func responseSize(url string) {
	// Schedule the call to WaitGroup's Done to tell goroutine is completed.
	defer wg.Done()

	fmt.Println("Step1: ", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Step2: ", url)
	defer response.Body.Close()

	fmt.Println("Step3: ", url)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Step4: ", len(body))
}

func main() {
	// Add a count of three, one for each goroutine.
	wg.Add(3)
	fmt.Println("Start Goroutines")

	go responseSize("https://www.golangprograms.com")
	go responseSize("https://stackoverflow.com")
	go responseSize("https://coderwall.com")

	// Wait for the goroutines to finish.
	wg.Wait()
	fmt.Println("Terminating Program")
}


// using mutex
package main

import (
	"fmt"
	"sync"
)

var (
	counter int32          // counter is a variable incremented by all goroutines.
	wg      sync.WaitGroup // wg is used to wait for the program to finish.
	mutex   sync.Mutex     // mutex is used to define a critical section of code.
)

func main() {
	wg.Add(3) // Add a count of two, one for each goroutine.

	go increment("Python")
	go increment("Go Programming Language")
	go increment("Java")

	wg.Wait() // Wait for the goroutines to finish.
	fmt.Println("Counter:", counter)

}

func increment(lang string) {
	defer wg.Done() // Schedule the call to Done to tell main we are done.

	for i := 0; i < 3; i++ {
		mutex.Lock()
		{
			fmt.Println(lang)
			counter++
		}
		mutex.Unlock()
	}
}

// fix fight for a single resource using atomic (sync/atomic) functions
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter int32          // counter is a variable incremented by all goroutines.
	wg      sync.WaitGroup // wg is used to wait for the program to finish.
)

func main() {
	wg.Add(3) // Add a count of two, one for each goroutine.

	go increment("Python")
	go increment("Java")
	go increment("Golang")

	wg.Wait() // Wait for the goroutines to finish.
	fmt.Println("Counter:", counter)

}

func increment(name string) {
	defer wg.Done() // Schedule the call to Done to tell main we are done.

	for range name {
		atomic.AddInt32(&counter, 1)
		runtime.Gosched() // Yield the thread and be placed back in queue.
	}
}

//pause play and stop goroutines
package main

import (
	"fmt"
	"sync"
	"time"
)

var i int

func work() {
	time.Sleep(250 * time.Millisecond)
	i++
	fmt.Println(i)
}

func routine(command <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	var status = "Play"
	for {
		select {
		case cmd := <-command:
			fmt.Println(cmd)
			switch cmd {
			case "Stop":
				return
			case "Pause":
				status = "Pause"
			default:
				status = "Play"
			}
		default:
			if status == "Play" {
				work()
			}
		}
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	command := make(chan string)
	go routine(command, &wg)

	time.Sleep(1 * time.Second)
	command <- "Pause"

	time.Sleep(1 * time.Second)
	command <- "Play"

	time.Sleep(1 * time.Second)
	command <- "Stop"

	wg.Wait()
}


##### channels 
// Simple program to demonstrate use of Buffered Channel

package main

import (
	"fmt"	
	"math/rand"
	"sync"
	"time"
)

var goRoutine sync.WaitGroup

func main(){
	rand.Seed(time.Now().Unix())

	// Create a buffered channel to manage the employee vs project load.
	projects := make(chan string,10)

	// Launch 5 goroutines to handle the projects.
	goRoutine.Add(5)
	for i :=1; i <= 5; i++ {
		go employee(projects, i)
	}

	for j :=1; j <= 10; j++ {
		projects <- fmt.Sprintf("Project :%d", j)
	}

	// Close the channel so the goroutines will quit	
	close(projects)
	goRoutine.Wait()
}

func employee(projects chan string, employee int) {
	defer goRoutine.Done()
	for {
		// Wait for project to be assigned.
		project, result := <-projects

		if result==false {
			// This means the channel is empty and closed.
			fmt.Printf("Employee : %d : Exit\n", employee)
			return
		}

		fmt.Printf("Employee : %d : Started   %s\n", employee, project)

		// Randomly wait to simulate work time.
		sleep := rand.Int63n(50)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		// Display time to wait
		fmt.Println("\nTime to sleep",sleep,"ms\n")

		// Display project completed by employee.
		fmt.Printf("Employee : %d : Completed %s\n", employee, project)
	}

}

#### Logging
// Program in GO language with real world example of logging.
package main

import (
"log"
"net/smtp"
)
func init(){
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("init started")
}
func main() {
// Connect to the remote SMTP server.
client, err := smtp.Dial("smtp.smail.com:25")
	if err != nil {	
		log.Fatalln(err)
	}
client.Data()
}
