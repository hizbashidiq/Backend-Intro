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

