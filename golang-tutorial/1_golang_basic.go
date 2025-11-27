// Clean Architecture -> one type of backend architecture pattern
// Delivery Layer  -> HTTP handlers, controller
// Use Cases       -> Application logic
// Domain/Entity   -> Business rule, core models
// Infrastructure  -> DB, external APIs, framework

// in clean architecture, dependencies always point inward
// handlers depen on use Cases
// use case depend on Domain
// domain depend on nothing
// so
// db cannot import use Cases
// use cases can import domain
// handlers can import use case

// Golang folder structure example
// /domain
//     user.go
// /usecase
//     create_user.go
// /infrastructure
//     postgresql_user_repo.go
// /delivery
//     http_handler.go

// not all architecture need to use clean architecture
// clean architecture usually used for
// medium to large systems
// systems that may change infrastructure i.e. change to mongodb from postgresql
// clean architecture make you think:
// if this framework disappear tomorrow, does my business code still work?
// can I test this logic without database running?


// CLI : Command Line Interface

// go mod init {project_name} -> initiate project that using Go modules, project name usually the same as folder name
// go run {file_name, ...}
// go test {file_name} -> testing package. file name need to end with _test.go
// go build -> create .exe file. defaulted to project_name. to change it, use the flag -o,
// go build -o {file_name.exe}
// go get -> to download package, similar like pip in Python (?). use flag -u to upgrade package
// use go get only in project directory
// go mod download -> similar with go get but doesn't change go.mod and go.sum. only download the one in go.mod
// usually used in CI/CD
// go mod tidy -> validating dependencies and downloading it if not downloaded yet
// go mod vendor




package main
// each project should have at least one package main. file with package main will executed first

import "fmt"
// import f "fmt" //I/O
// f is an alias, so you can just write f.Println()
// import "strings"
import (
	"math"
	"strings"
  . "golang-tutorial/library" // root folder is where go.mod is
  // use . to let golang think that it's in the same level so you don't need to write library.Student, just Student
  // but it's not recommended (?)
  "reflect" //to inspect variable, metadata of variable: structure, tipe, value, pointer, etc
  // also to modify unknown data type
  "runtime"
  "time"
  "math/rand"
  "strconv"
  "errors"
  "os"
  "regexp"
  "encoding/base64"
  "crypto/sha1"
  "flag"
  "os/exec"
  "io"
)

// from "fmt" import Println, we can't do that

// import "fmt/Println"

// import f "fmt" -> alias

// so since you need to use all package imported in Go, is using blank import the best practice? (NOPE)

func main() {
	fmt.Println("Hello from Go in VS Code!")
	var firstName string = "John" // use this outside function i.e. create a package
	// you can also leave the type out i.e. var firstname = "John"
	var lastName string
	lastName = "Wick"

	// fmt.Println("Hello %s %s !", firstName, lastName)
	fmt.Printf("halo %s %s!", firstName, lastName) // need to use printf (not println) to use %s
	// %s string, %d int, %f float, %v any value
	fmt.Println("halo", firstName, lastName) //auto add space
	fmt.Println("halo "+firstName+ " "+lastName)

	middleName := "Baba Yaga" // use this inside function

	// middleName := "The Boogeyman" can't reassign value this way, unless there's new variable for example

	x, middleName := 1, "The Boogeyman"
	x += 1

	// you can assign with =
	middleName = "Jardani Jovonovich"

	fmt.Printf("Hello %s %s %s!\n", firstName, middleName, lastName)

	// underscore variable, to store unused variable
	// _ := "hewwo", this is error since underscore isn't considered new
	_ = "hewwo"
	// also you can't use the value again
	// fmt.Printf("hello %s", _)

	// new() is to store pointer
	p := new(string)
	fmt.Println(p)
	fmt.Println(*p)

	// make() used to declare slice, channel, map

	decimalVar := 2.34
	fmt.Printf("decimal %f\n", decimalVar)
	fmt.Printf("decimal 3 %.3f\n", decimalVar)

	intVar := 234
	fmt.Printf("integer %04d\n", intVar)

	var exist bool = false
	fmt.Printf("exist? %t \n", exist)

	// Go philosopy: Be specific when you care, be generic when you don't

	// to create long or complex string use `` (backtick)

	const pi = 3.14

	// msg := "hello ", pi
	msg := fmt.Sprintf("Hello %f", pi)

	fmt.Printf("%s \n", msg)
	var(
		one = true
		two = "two"
		three = 3
	)
	fmt.Printf("%t %s %d\n", one, two, three)

	const four, five = 4, "five"
	fmt.Println(four, five)

	// and -> &&, or -> ||

	// can't do ternary in Go
	// only if, else if, else and switch

	// temporary variable -> only in if, switch, and for statement
	var number = 7400.0

	if percent := number/100; percent >= 80 {
		fmt.Printf("Since your number is %.2f your grade is A", percent)
	} else if percent >= 60 && percent < 80 {
		fmt.Printf("Since your number is %.2f your grade is B\n", percent)
	} else if percent >=40 && percent < 60 {
		fmt.Printf("Since your number is %.2f your grade is C\n", percent)
	} else {
		fmt.Printf("Since your number is %.2f your grade is D\n", percent)
	}

	switch number{
	case 80:
		fmt.Printf("Since your number is %.2f your grade is A\n", number)
	default:
		fmt.Printf("Since your number is %.2f your grade is E\n", number)
	}

	switch percent := number/100;{
	// case percent := number/100; percent > 80:
	case percent > 80:{
		fmt.Printf("Since your number is %.2f your grade is A\n", number)
		fmt.Printf("You're great")
	}
	default:{
		fmt.Printf("Since your number is %.2f your grade is E\n", number)
		fmt.Println("You need to learn more")
	}
	}
	// use fallthrough in switch if you need to keep checking condition

	for i := 0; i < 5; i++{
		fmt.Println("Angka", i)
	}

	// while
	var j = 0
	for j < 5{
		fmt.Println("Nomor", j)
		j++
	}

	test := "abcdefgh"

	for i, v := range test{
		fmt.Println("Index", i, "value", v)
	}

	for i, v := range test{
		fmt.Printf("Index %d value %c \n", i, v)
	}
	myArray := [5]int{1,3,5,7,9}

	for _, v:= range myArray{
		fmt.Println("value", v)
	}
	sliceMyArray := myArray[0:2]

	for _, v:= range sliceMyArray{
		fmt.Println("value", v)
	}

	myMap := map[byte]int{'a':0,'b':1,'c':2,'d':3}

	// map is unordered
	for k,v:=range myMap{
		fmt.Println("key", k, "value", v)
		fmt.Printf("key %c value %d \n", k, v)
	}

	// break used to stop a loop, continue used to skip 1 iteration

	outerLoop: //labelling
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
			break outerLoop
			// break
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}

	var names [10]string
	for i:= range names{
		names[i] = fmt.Sprintf("number %d", i)
	}
	for i:= range names{
		fmt.Println(names[i])
	}

	fmt.Println(len(names))
	fmt.Println(names)

	var num_arr = [...]int{1,2,3,4,5,6,7}
	fmt.Println(num_arr)
	// can't append array, only slice
	// slice is dynamic while array is static
	// that's why 99% you will use slice
	// and slice is kinda a pointer(?), so you can make an s1 := s2, you change s2, s1 also changed
	// since it's refer to the same address
	// use s2 = make([]int, len(s1)), copy(s2, s1) so it's refer to different memory address
	// wait so it kinda depends of cap, if you append s2, it possible the s1 didn't change because
	// maybe the cap is full so they need to reallocate memory for s2 but s1 remain the same (?)

	var matriks [][]int
	matriks = append(matriks, []int{1,2,3})
	matriks = append(matriks, []int{4,5,6})
	fmt.Println(matriks)

	var fruits = make([]string, 2)
	fruits [0] = "banana"
	fruits [1] = "strawberry"

	// cap() is confusing. and I don't know the usage of it lol
	// make([]int, len, cap) i.e. make([]int, 3, 5)


	// append()
	// so if cap() == len(), they reallocate
	// else (cap>len) they keep continuing the memory address
	var animals = []string{"elephant","cat","dog"}
	xanimals := animals [:2] //cap = 3, len = 2
	yanimals := append(xanimals, "girrafe") //change dog into girrafe in animals
	zanimals := append(yanimals, "whale")
	zanimals[2] = "crow"
	fmt.Println(animals)
	fmt.Println(xanimals)
	fmt.Println(yanimals)
	fmt.Println(zanimals)

	// aanimals := animals[start_from:end_before:cap()]

	var map1 map[string]int
	map1 = map[string]int{} // the {} is important, if you don't include it it's gonna assign nil so you can't change the value
	// map1 = make(map[string]int)
	// map1 = *new(map[string]int)
	map1["january"] = 1
	map1["february"] = 2

	// delete(map1, "january")

	var value, isExist = map1["march"] // isExist gonna be true/false regarding is it exist or not
	fmt.Println(value)
	fmt.Println(isExist)

	array1 := []string{"John","Wick","The","Boogeyman"}
	msg = "Hello"

	printMessage(msg, array1)

  fmt.Println(circle(15))
  fmt.Println(average(1,2,3,4,5,6))

  array2 := []int{1,2,3,4,5,6,7,8,9}
  fmt.Println(average(array2...))

  // closure function can carry state, and keep things private

  // pointer


  // & is used to referencing -> x = 5 -> &x = memory address
  // * is used to dereferencing ->
  var int1 int = 10
  var string1 string = "hello"
  fmt.Println(int1, string1)
  fmt.Println(&int1, &string1)

  var p_int1 *int = &int1
  var p_string1 *string = &string1
  fmt.Println(p_int1, p_string1)
  fmt.Println(*p_int1, *p_string1)

  // Go don't use Class, instead it use Struct

  var student1 student
  student1.name = "John"
  student1.grade = 6

  var student2 = student{}
  var student3 = student{grade: 4, name: "Doe"}
  fmt.Println(student1)
  fmt.Println(student2)
  fmt.Println(student3)

  var s1 *student = &student1
  fmt.Println(s1)
  fmt.Println(*s1)
  s1.name = "Wick"
  fmt.Println(*s1)
  fmt.Println(s1)
  fmt.Println(&s1)
  fmt.Println(student1)
  fmt.Println(&student1)

  // embedded struct, so you don't need to write s.person.age, only s.age as long as there's no name conflict
  var s5 = students{}
  s5.name = "wick"
  s5.age = 21
  s5.grade = 2

  fmt.Println("name  :", s5.name)
  fmt.Println("age   :", s5.age)
  fmt.Println("age   :", s5.person.age)
  fmt.Println("grade :", s5.grade)
  // anonymous struct
  var s6 = struct{
    person
    grade int
  }{} // this is important to initiate zero value to the property
  fmt.Println(s6)
  var s7 = []struct{
    person
    grade int
  }{}
  fmt.Println(s7)

  var s8 struct{
    person
    grade int
  }
  fmt.Println(s8)

  // nested struct
  type owner struct{
    identification struct{
      id int
      name string
    }
    car_name string
  }
  var o1 owner
  fmt.Println(o1)

  // you can also do it horizontaly using ; to make a struct

  // tag is important later on, usually used for json, database, and validation
  s1.sayHello()

  // you need to make method pointer if you want to change a property value (?)
  // you only change the value inside of method otherwise

  // in Go, 1 folder = 1 package
  // first letter upper case -> exported/public
  // first letter lower case -> unexported/private
  // that's applicable to all (function, struct, method, property)

  // library.SayHello("John")
  // library.introduce("John")

  // var s9 = library.student{"Ethan",21}
  // var s9 = library.Student{"Ethan", 21}
  // s9 := library.Student{"Ethan", 21}
  s9 := Student{"Ethan", 21}
  // var s9 library.student = {"Ethan", 21}

  var s10 string = "hewo"

  fmt.Println(s9, s10)

  // since sayHola is in same package but different file, just add it in the command
  // go run 1_golang_basic.go partial.go
  // or better -> go run *.go
  sayHola()
  // since partial.go is the same package (main package) you can access private/unexported things

  // you can use more than 1 init() in a file
  // the order of init() execution are random between files
  // but it'll be in order of import, so if it's same package it could be random (?)

  var bangunDatar hitung
  bangunDatar = persegi{10.0}
  fmt.Println(bangunDatar.luas(), bangunDatar.keliling())

  bangunDatar = lingkaran{14}
  fmt.Println(bangunDatar.(lingkaran).jariJari()) // what is this?
  // oh so this is casting bangunDatar to lingkaran first so you can use jariJari method
  // not casting, it's type assertion and only works for interface
  fmt.Println(bangunDatar.luas(), bangunDatar.keliling())

  // var int2 = 20
  // fmt.Println(int2)
  // fmt.Println(int2.(float64))

  // pointer dereference are slower than straight value in small and simple value (int, string, etc)
  // if you don't want the function to change anything about parameter, use value instead of pointer (mutability)
  // pointer can complicate concurrency

  // so only use pointer if:
  // 1. you need to modify the parameter value
  // 2. the struct is large because it's cheaper to call/copy
  // if you have a pointer method (usually so you can modify?) and a value method in a struct, you can access
  // both using a pointer variable but can't access pointer method using value variable
  // the most important rule of thumb is using pointer if you want to modify (?)

  var secret interface{} //empty interface can be use for any datatype
  secret = "hello"
  fmt.Println(secret)
  secret = 20
  fmt.Println(secret)

  // empty interface is a datatype btw, not an object
  // any is alias of interface{}
  // only in 1.18+ version btw
  // var secret2 any
  // secret2 = "test"
  // fmt.Println(secret2)
  // secret2 = 40
  // fmt.Println(secret2)

  fmt.Printf("%T\n", secret)
  // var number2 = secret * 2 //since secret is an interface this is error, you need to type assertion first
  var number2 = secret.(int)*2
  fmt.Println(number2)

  secret = []string {"apple", "banana", "cherry"}
  // var f_fruits = strings.Join(secret, ", ") -> error, need to type assertion
  var f_fruits = strings.Join(secret.([]string), ", ")
  fmt.Println(secret)
  fmt.Println(f_fruits)

  // & -> get the address of this value, used when you want to create pointer
  // * has two meanings
  // 1. in variable declaration -> defines a pointer type (x *int)
  // 2. when reading a value -> dereference, get the actual value stored at the pointer (*x = 100)

  type person20 struct{
    name string
    age int
  }
  // var secret3 interface{} = &person{name: "Wood", age:24}
  secret3 := interface{}(&person{name: "Wood", age:24})

  fmt.Println(secret3)
  // fmt.Println(*secret3) error because * can only be applied to a pointer value, but secret3 is an interface data type
  fmt.Println(*secret3.(*person))
  fmt.Printf("%p",*secret3.(*person))
  // fmt.Printf("%p",&secret3.(*person))
  fmt.Printf("%p\n", &secret3)

  // p := s.(*person)
  // fmt.Printf("%p\n", p)      // address stored in interface
  // fmt.Printf("%p\n", &p)     // address of the pointer variable 'p'
  // fmt.Printf("%p\n", &(*p))  // address of the underlying struct
  // yeah I'm confused

  // fmt.Println("%p", secret3)
  // fmt.Println("%p", *secret3)
  var name2 string = secret3.(*person).name
  fmt.Println(name2)

  // reflect.ValueOf -> return reflect.Value
  // reflect.TypeOf -> return reflect.Type
  // usually used to work with unknown data types at run time
  // so it should be together with interface{} since it can be any type
  // also to work with framework-level code so it can receive any data type
  var number3 = 3
  reflectVal := reflect.ValueOf(number3)
  fmt.Println(reflectVal.Type())

  if reflectVal.Kind() == reflect.Int{
    fmt.Println(reflectVal.Int())
  }
  // Kind() broad type i.e. int, string, slice
  // always return most outer like []map[int]string gonna omit slice i.e.[]
  // to go inner you'll use .Elem().Kind()-> map, Key().Kind()->int Elem().Kind() -> string
  // Type() full information i.e. map[string]int, main.Person
  // ValueOf() : normal -> reflect.Value
  // Interface() : reflect.Value -> normal
  // reflect can access information of all kind of things (variable, method, struct) but it needs to be public/exported

  // goroutines are like minithread. it's asynchronous -> concurrent programming in Golang
  // concurrency != paralel != multihreading
  // concurrency means you don't wait for the result of other hardware
  //  and continue living and go back when result is done



  runtime.GOMAXPROCS(2) // to tell how much core gonna be used
  go print(5, "halo")
  // print(5, "hola")
  print(5, "what's that? Ugh")

  // var input string
  // fmt.Scanln(&input)
  fmt.Println(runtime.NumCPU())

  // channel -> to connect between goroutines by send-receive data mechanism
  // send-receive in channel is synchronous/blocking
  runtime.GOMAXPROCS(2)

  var messages = make(chan string)

  var sayHelloTo = func(who string){
    var data = fmt.Sprintf("Hello %s", who)
    messages <- data
  }
  go sayHelloTo("John")
  go sayHelloTo("Wick")
  go sayHelloTo("Baba")

  var message1 = <- messages
  var message2 = <- messages
  var message3 = <- messages // this just means that current goroutines tries to receive a value from channel
  fmt.Println(message1, message2, message3)

  // channel can be a parameter i.e. func f(messages chan string)
  // channel in parameter is pass by reference i.e. pointer, not a value

  // IIFE (Immediately Invoked Function Expression) ->anonymous function that directly called when declared

  // in default data transfer in channel is unbuffered
  // buffered channel use to make it asynchronous as long as n_data <= n_buffer
  // if buffer full, it'll wait until buffer not full
  // in buffered channel, send data is asynchronous, receive data is synchronous
  // receive data is synchronous means if channel empty receive blocks and wait
  // like the receive can't do anything (?)
  // channel is maibox (not sender, nor receiver)
  // main() is goroutine


  runtime.GOMAXPROCS(2) // tell how much core can be used for parallel goroutine, default at max i.e. runtime.NumCPU()
  m := make(chan int, 3) //it's start from 0, so 3 buffer means it has 4 capacity (kinda weird tbh)
  go func(){
    for{ //while true
      i := <- m
      fmt.Println("Receive data", i)
    }
  }()

  for i:=0;i<5;i++{
    m <- i
    fmt.Println("Send data", i)

  }
  time.Sleep(1*time.Millisecond) // pause current goroutine i.e. main() usually
  // in real code wg used instead of time.sleep
  // range ch in for loop already use <- ch, yeah idk
  // range ch also automatically closed when channel closed
  // for v:= range ch is like v, ok := <-ch with v is value and ok is isChannelOpen

  // channel is to sharing data and communication between goroutines
  // not to control goroutine execution

  runtime.GOMAXPROCS(2)
  arr := []float64{1.2,2.4,3.5,5.6,9,2}
  // var ch1 chan float64 can't declare like this since it's initiate nil channel so it'll deadlock
  ch1 := make(chan float64)
  go avg(arr, ch1)
  // var ch2 chan float64
  ch2 := make(chan float64)
  go findMax(arr, ch2)

  for i:=0;i<2;i++{
    select{
    case aveg := <- ch1:
      fmt.Println(aveg)
    case max := <- ch2:
      fmt.Println(max)
    }
  }
  // select is strictly for channel
  // the cases either send to channel (ch<-v), receive from channel (v:=<-ch), waiting on channel close, or default

  ch3 := make(chan int)
  go sendMessages(ch3)
  // go printMessages(ch3)
  printMessages(ch3)

  // channel direction
  // ch chan string : both send and receive
  // ch chan<- string : just send
  // ch <-chan string : just receive
  runtime.GOMAXPROCS(2)
  var ch4 = make(chan int)
  go sendData(ch4)
  retrieveData(ch4)

  // defer is to execute a program before a function finish, if >1 defer, it'll executed sequentially
  // os.Exit() force exit program immediately

  // go use Error,Panic, and Recover not Exception
  // error is a datatype. it has Error() method that return detail of error

  defer catch()

  var input string
  // fmt.Printf("Enter a number: ")
  // fmt.Scanln(&input)
  input = "12345"
  var number4 int
  var err error
  number4, err = strconv.Atoi(input)

  if err == nil{
    fmt.Println(number4, "is a number")
  } else{
    fmt.Println(input, "is not a number")
    fmt.Println(err.Error())
  }

  if value1, err1 := validate(input); value1{
    fmt.Println("Your input is", input)
  }else{
    fmt.Println("Your input is empty")
    fmt.Println(err1)
    // btw print err and print err.Error() are the exact same thing
    // but err has strings function modification so it print as such
    // if you want to pass the message as a string, then use err.Error() since it's a string
    fmt.Printf("%T %T\n", err1, err1.Error())
    // panic(err1.Error())
    panic(err1)
    // panic(err.Error()) and panic(err) also exactly the same, just use panic(err)

    // if there's panic, deter that is above panic will be executed first
    fmt.Println("hafdoafhds")
  }
  fmt.Println("Nggoeheee")

  // so recover isn't try catch thing, it's just let panic be in that function and continue in outer function
  // but since this is only 1 function, the print hafdoafhds will never run because it's still in the same function
  // where the recover happen
  // String Format Layout
  // %b -> biner
  // %c -> character, or from int to unicode
  // %d -> int base 10
  // %e or %E -> scientific notation
  // %f or %F -> float, %.nf where n is how much number behind decimal
  // %g -> float but a lot of decimal number
  // %o -> string numeric base 8
  // %p -> address, used with &value(needed?)
  // %q -> escape string
  // %s -> string
  // %t -> bool
  // %T -> datatype
  // %v -> anything, used for debug usually, only value, %+v with the name i.e. name:wick, while %v only wick
  // %#v -> more complete at what package, and how it's declared
  // %x or %X -> numeric to string base 16
  // %% -> to print "%"

  // Random Number Generator (RNG)
  // math/rand is a pseudo RNG
  randomizer1 := rand.New(rand.NewSource(10))
  for i:=0;i<3;i++{
    fmt.Println(randomizer1.Int())
  }
  // Go has no special method like __str__ or __init__, New() are just a convention
  // but Go has String() method that return fmt.Sprintf, fmt uses it if available in a struct

  now := time.Now().UTC().UnixNano()
  randomizer2 := rand.New(rand.NewSource(now))
  fmt.Println(now)
  for i:=0;i<3;i++{
    fmt.Println(randomizer2.Int())
  }

  // randomizer2.Intn(100) give number range 0-99
  // for random string you need to create your own function? well

  time1 := time.Now()
  time2 := time.Date(2011,12,24,10,20,0,0, time.UTC)
  fmt.Printf("%v\n%v\n", time1, time2)
  // some time.Time method time1.Year(), .Month(), etc
  // Go is weird it use 2006 as year, 006 as 3 digit year, 06 as 2 digit year, 02 as 2 digit day, 01 as 2 digit month, etc
  // some predefined layout format for parsing time is time.RFC822

  var date1, _ = time.Parse(time.RFC822, "02 Sep 15 08:00 WIB")
  // Mon Jan _2 15:04:05 MST 2006 UnixDate
  var date2, _ = time.Parse(time.UnixDate, "Sat Sep 02 08:00:00 WIB 2006")

  fmt.Println(date1.String())
  fmt.Println(date2.String())

  // from time.Time to string
  var date1S1 = date1.Format(time.RFC822)
  var date1S2 = date2.Format(time.UnixDate)
  fmt.Println(date1S1)
  fmt.Println(date1S2)


  // fmt.Println(i)
  // i think it's return date and error so the _ is error thing

  var timer = time.NewTimer(1 * time.Millisecond)
  fmt.Println("start")
  <-timer.C
  fmt.Println("finish")

  var ch5 = make(chan bool)
  time.AfterFunc(1 * time.Millisecond, func (){
    fmt.Println("Expired")
    ch5 <- true
  })
  fmt.Println("Start ch5")
  xx := <-ch5
  fmt.Println("Finish ch5")
  fmt.Println(xx)

  // time.Sleep() and time.After() is similar, but time.After() will return channel data so you need <-
  <-time.After(1 * time.Millisecond)
  // so time.After() can be used as case in select, and it's kinda the main reason for it I guess

  // ch6 := make(chan bool)
  // ticker := time.NewTicker(time.Millisecond) //everysecond ticker will send date information via .C
  // // ticker.C is a channel

  // go func(){
  //   time.Sleep(1 * time.Millisecond)
  //   ch6 <- true
  // }()

  // for {
  //   select{
  //   case <-ch6:
  //     ticker.Stop()
  //     return //terminate goroutine and main()?
  //   case t:=<-ticker.C:
  //     fmt.Println("Hello", t)
  //   }
  // }
  fmt.Println("fadsjkldajshlasdfk")
  var timeout = 60
  ch := make(chan bool)

  go timer2(timeout, ch)
  go watcher(timeout, ch)

  var input2 string
  input2 = "dafsafd"
  // fmt.Scanln(&input2)
  fmt.Println("good", input2)
  time.Sleep(1*time.Millisecond)

  start2 := time.Now()
  // start2 := "hewwo"
  time.Sleep(1*time.Millisecond)
  // duration := time.Since(time2)
  duration := time.Since(start2)


  fmt.Println(duration.Seconds(), duration.Minutes(), duration.Hours())

  start3 := time.Now()
  duration2 := start3.Sub(start2)
  fmt.Println(duration2.Seconds(), duration2.Minutes(), duration2.Hours())

  // okay so time.nano = 1, and time.second is a constant = 1000000000, so you can do time.Sleep(1) for 1 nanosecond

  // var n = 5
  // dur := n * time.Second error because time.Second is time.duration while n is int
  // so you need to declare var n time.duration = 5 or casting dur := time.duration(n) * time.Second

  // strconv
  // Atoi is ASCII to Int okay
  // so Itoa is Int to ASCII
  // .ParseInt(str, base, 64) 64 here is int64, I suppose you can change it to other, but idk other option
  // -> exactly the same with Atoi, just more control
  // .FormatInt(num, base) -> int64 to string
  // .ParseFloat() -> string to float
  // etc

  // strings.Contains(p1, p2) is p2 part of p1
  // strings.HasPrefix(p1, p2) is p2 a start of p1
  // strings.HasSuffix(p1, p2) is p2 an end of p1
  // strings.Count(p1, p2) p2 is a character (?), count how much p2 is used in p1
  // etc

  // RegEx in Go is using RE2

  var text = "banana burger soup"
  var regex, err2 = regexp.Compile(`[a-z]+`)

  if err2 != nil{
    fmt.Println(err2)
  }
  var res1 = regex.FindAllString(text, -1)
  fmt.Printf("%#v\n", res1)

  var s11 = "John Wick"

  var encodedString = base64.StdEncoding.EncodeToString([]byte(s11))
  var decodedByte, _ = base64.StdEncoding.DecodeString(encodedString)
  var decodedString = string(decodedByte)
  fmt.Println(encodedString)
  fmt.Println(decodedByte)
  fmt.Println(decodedString)

  // .Encode more memory efficient since it'll only use len(data) than EncodeToString
  // if you encode 200MB file, EncodeToString will make 266MB string

  var encoded = make([]byte, base64.StdEncoding.EncodedLen(len(s11))) //return how much you memory you need
  // EncodedLen always allocate bigger than (len(s11)) using the calculation
  base64.StdEncoding.Encode(encoded, []byte(s11)) //so encoded = how much you allocate memory
  encodedString = string(encoded)

  var decoded = make([]byte, base64.StdEncoding.DecodedLen(len(encoded)))
  // while DecodedLen allocate <= len(encoded) using inverse calculation in EncodedLen
  var _, err10 = base64.StdEncoding.Decode(decoded, encoded)
  if err10 != nil{
    fmt.Println(err)
  }
  decodedString = string(decoded)

  fmt.Println(encoded)
  fmt.Println(encodedString)
  fmt.Println(decoded)
  fmt.Println(decodedString)

  // for URL, use URLEncoding
  var url = "https://kalipare.com"
  encodedString = base64.URLEncoding.EncodeToString([]byte(url))
  decodedByte, _ = base64.URLEncoding.DecodeString(encodedString)
  decodedString = string(decodedByte)
  fmt.Println(url)
  fmt.Println(encodedString)
  fmt.Println(decodedByte)
  fmt.Println(decodedString)

  // SHA1 -> Secure Hash Algorithm 1
  // length of SHA1 is 160 bit, 20 byte, usually presented in 40 hexa
  var text1 = "this is secret"
  var sha = sha1.New()
  sha.Write([]byte(text1)) //set data that gonna be hashed. it should be in []byte
  encrypt := sha.Sum(nil)

  fmt.Printf("%T %T\n", sha, encrypt)
  fmt.Println(encrypt)
  fmt.Printf("%x\n", encrypt)
  fmt.Printf("%s\n", encrypt)

  // wait so you don't need to use Write if you pass parameter directly to Sum([]byte(text1))
  // okay so use sha1.New()+Write() if the data large and or incomplete, else just use sha1.Sum
  var text2 = "this is not a secret"
  encrypt2 := sha1.Sum([]byte(text2))
  encrypt2String := fmt.Sprintf("%x", encrypt2)
  fmt.Println(encrypt2)
  fmt.Printf("%x\n", encrypt2)
  fmt.Println(encrypt2String)

  var args = os.Args // path + args
  // to get only args just go args[1:]
  fmt.Println(args)
  fmt.Printf("%#v\n", args)

  // my guess is it's (flag_name, default_value, help) yep i'm right
  var name4 = flag.String("name4", "user", "type your name") // return string pointer so need to dereference to *name
  var age4 = flag.Int64("age4", 18, "type your age")

  // flag.Parse()
  // fmt.Printf("%s %d\n", name, age)


  // you can use -name="name" or -name "name"
  // use Var after flag.String (or any type) so you can use variable directly without dereference
  var name5 string
  var age5 int64

  flag.StringVar(&name5, "name5", "user", "type your name2")
  flag.Int64Var(&age5, "age5", 18, "type your age2")

  flag.Parse() //IMPORTANT BUT ONLY ALLOWED ONCE

  fmt.Printf("%s %d\n", *name4, *age4)
  fmt.Printf("%s %d\n", name5, age5)

  // {program_path} --help -> to write the help thing

  var output1, _ = exec.Command("ls").Output()
  var output2, _ = exec.Command("pwd").Output()
  var output3, _ = exec.Command("git", "config", "user.name").Output()

  fmt.Printf("%T %T %T\n", output1, output2, output3)
  // uint8 is alias for byte
  fmt.Println(output1)
  fmt.Println(output2)
  fmt.Println(output3)
  fmt.Printf(" -> ls\n%s\n", string(output1))
  fmt.Printf(" -> pwd\n%s\n", string(output2))
  fmt.Printf(" -> git config user.name\n%s\n", string(output3))
  fmt.Println("int 100 string ", string(int(100)))

  fmt.Println(runtime.GOOS)
  // use runtime.GOOS == "windows" else (linux/macOS) for OS specific command

  path, _ := os.Getwd()
  path = path+"/files/test.txt"
  // path = path+"/1_golang_basic.go"

  fmt.Println(path)
  // createFile(path)
  writeFile(path)
  readFile(path)

  // Damn, so there's shadowing rule i.e. declare same variable in different scope

  // delete file
  // os.Remove(path)

}

func readFile(path string) error{
  var file, err = os.OpenFile(path, os.O_RDONLY, 0644)
  if err != nil {return err}
  defer file.Close()

  var text = make([]byte, 1024) //1024 is = 1KB
  for {
    n, err := file.Read(text)
    fmt.Println(err)
    fmt.Println(n)
    if err == io.EOF{
      break
    }
    if err != nil{
      return err
    }
    if n == 0{
      break
    }
  }
  fmt.Println(string(text))
  // fmt.Println(string(text[2]))

  fmt.Println("File has been read!")
  return nil
}

func writeFile(path string) error{
  // var file, err = os.OpenFile(path, os.O_RDWR, 0644)
  var file, err = os.OpenFile(path, os.O_APPEND, 0644)

  if err != nil {return err}

  defer file.Close()

  if _, err := file.WriteString("Hello\n"); err != nil{
    return err
  }

  if _, err := file.WriteString("I mean HEWWO\n"); err != nil {
    return err
  }

  // save
  // file.Sync() are not needed. if it's a crucial data, then that's different story
  // if err := file.Sync(); err != nil {
  //   return err
  // }

  fmt.Println("File has been edited!")
  return nil
} //this is erase all the data btw, not append the data

func createFile(path string) error{
  if _, err := os.Stat(path); os.IsNotExist(err){
    file, err := os.Create(path)

    if err != nil{
      return err
    }
    defer file.Close() // you don't need defer here since it's just creating file, but it's best practice for bigger code
    fmt.Println("File created successfully!")
  }
  return nil
}

func timer2(timeout int, ch chan<- bool){
  time.AfterFunc(time.Duration(timeout) * time.Second, func(){
    ch<-true
  })
}

func watcher(timeout int, ch <-chan bool){
  <-ch
  fmt.Println("\ntime out, no answer more than", timeout, "seconds")
  os.Exit(0)
}

func catch(){
  if r:=recover();r!=nil{
    fmt.Println("Error occured", r)
  }else{
    fmt.Println("Program running succesfully")
  }
}

func validate(input string) (bool, error){
  if strings.TrimSpace(input) == ""{
    return false, errors.New("Cannot be empty")
  }else{
    return true, nil
  }
}

func sendData(ch chan<- int){
  randomizer := rand.New(rand.NewSource(time.Now().Unix()))

  for i:=0;true;i++{
    ch <- i
    time.Sleep(time.Duration(randomizer.Int()%10+1) * time.Millisecond)
  }
}
func retrieveData(ch <-chan int){
  loop:
  for{
    select{
    case data := <- ch:
      fmt.Println("hewwo")
      fmt.Println(data)
    case <-time.After(1 * time.Millisecond):
      fmt.Println("Timeout. No activities after 1 seconds")
      break loop
    }
  }
}

// func sendMessages(ch chan int){
func sendMessages(ch chan<- int){
  for i:=0;i<20;i++{
    ch <- i
  }
  close(ch)
}
// func printMessages(ch chan int){
func printMessages(ch <-chan int){
  for message := range ch{ //keep receiving until ch closed
    fmt.Println(message)
  }
}
func avg(arr []float64, ch chan float64){
  sum := 0.0
  for _,v:=range arr{
    sum += v
  }
  ch <- sum/float64(len(arr))
}
func findMax(arr []float64, ch chan float64){
  max := arr[0]
  for _,v:=range arr{
    if max<v{
      max = v
    }
  }
  ch <- max
}


func print(till int, message string){
  for i:=0; i<till; i++{
    fmt.Println((i+1), message)
  }
}
// interface is a data type
// you need all method in an interface to be able to declared as said interface
// but you can have other method outside of the interface
// for example, in hitung you have luas and keliling and lingkaran you have both + jarijari
// it's still valid to lingkaran be that interface data type
// but if you have only luas, you can't be assign to hitung interface variable
// you can embed interface to another interface, like if you create hitung3d, you can have a member of hitung
type hitung interface{
    luas() float64
    keliling() float64
}

type lingkaran struct{
    diameter float64
}
func (l lingkaran) jariJari() float64{
    return l.diameter/2
}
func (l lingkaran) luas() float64{
    return math.Pi*math.Pow(l.jariJari(),2)
}
func (l lingkaran) keliling() float64{
    return math.Pi*l.diameter
}

type persegi struct{
    sisi float64
}
func (p persegi) luas() float64{
    return math.Pow(p.sisi,2)
}
func (p persegi) keliling() float64{
    return p.sisi*4
}

type student struct{
  name string
  grade int
}

func (s student) sayHello(){
  fmt.Printf("Hello, my name is %s. I'm %d grade.", s.name, s.grade)
}

type person struct {
  name string
  age  int
}

type students struct {
  grade int
  person
}

/*
	multi line comment
*/
func printMessage(msg string, arr []string){
	// for i:= range arr{
	// 	msg = fmt.Sprintf("%s %s",msg, arr[i])
	// 	fmt.Println(arr[i])
	// }
	var nameString = strings.Join(arr, " ")
	fmt.Println(msg, nameString)
}

// func calculate(d float64) (float64, float64){
func circle(d float64) (area float64, circumference float64){
  // area := math.Pi * math.Pow(d/2,2)
  // circumference := math.Pi * d
  // return area, circumference
  area = math.Pi * math.Pow(d/2,2)
  circumference = math.Pi * d
  return
}

// variadic function
func average(numbers ...int)(avg float64){
  var total int
  for i:=range numbers{
    total += numbers[i]
    // fmt.Println(total)
  }
  // fmt.Println(len(numbers))
  // int/int -> return int in Golang
  avg = float64(total)/float64(len(numbers))

  return
}

