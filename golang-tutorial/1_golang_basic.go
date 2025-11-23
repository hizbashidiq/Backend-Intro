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


import "fmt" //I/O

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

}

/*
	multi line comment
*/
