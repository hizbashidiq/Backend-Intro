package main

// CLI
// GIT

// COMPUTATIONAL THINKING -> a way of thinking logically and solving problems in an organized manner

// 4 cornerstones
// 1. Decomposition -> breaking down a complex problem into smaller, more menageable parts
// 2. Pattern recognition -> looking for similarities among and within problems
// 3. Abstraction -> focusing on important part only, ignoring irrelevant detail
// 4. Algorithms -> developing a step by step solution to a problem, or the rules to follow to solve a problem

// BASIC PROGRAMMING IN GO
// -statically typed
// -strong emphasis on conccurrency
// byte is the alias for uint8
// rune is the alias for int32

import (
	"fmt"
	"sync"
	// "math"
	// "strings"
	// "sort"
	"errors"
)

func main(){
  var n1 float64 = 1.2
  var n2 float64 = 1
  fmt.Printf("%.10f\n",n1-n2)

  var i int32 = 10
  j := i
  var k *int32 = &i
  fmt.Println(j, &j)
  fmt.Println(i, &i)

  fmt.Println(*k, k, &k)
  *k = 5 //*k is dereference of k, which stores i address, that's why this change i value
  // *(&x) = x -> &k is an address, * is dereference
  fmt.Println(i)

  // avoid magic number (hardcoded value that may appear in the code without context)
  // i.e. for i:=0;i<4;i++ -> 4 is a magic number it's better to store 4 in a variable/constant first

  // bitwise operator (&, |, ^, <<, >>)
  var l int = 10
  m := int32(l)
  fmt.Printf("l %T\n", l)
  fmt.Printf("m %T\n", m)

  // pointer mainly used if we want to change a variable within a struct, function or method
  // less common but pointer also used to avoid copy big struct (performance), just use the address
  // and some other usage

  // pass by value most of the time gonna copy the object/data, meanwhile pass by reference only copy the address,
  // but still referencing to same object/data

  // slice hold a reference to underlaying array
  // that's why if you assign a slice to a variable, both share the same object

  myMap := map[string]int{"test0":0,"test1":1}
  // check if a key exist
  _,ok := myMap["test0"]
  fmt.Println(ok)
  _,ok = myMap["test2"]
  fmt.Println(ok)

  // if we have a struct pointer, Go allow to implicit dereferencing so you can just use it as usual
  p := &MyStruct{10}
  fmt.Println(p.i) // you don't need to write *p.i
  fmt.Println(&p) // address of the struct i.e. MyStruct
  fmt.Printf("%p\n", p) // address of p

  // if a function only used once or a few times, it might be a good candidate to be an anonymous function
  // closure is an anonymous function that capture a variable within a scope basically(?)
  // so it can use or change the variable even after the outer execution is complete

  // defer are LIFO

  // Varible Hands On
  products := []variable{
    {"A", 100000, 200, 0},
    {"B", 67000, 12, 20},
    {"C", 56000, 80, 0},
    {"D", 1000, 1350, 0},
    {"E", 20000, 1, 0},
    {"F", 38455, 7, 15},
    {"G", 76000, 5644, 0},
    {"H", 530120, 30, 10},
    {"I", 143000, 54, 0},
    {"J", 16000, 109, 0},
  }

  var annualGrossProfit int

  for _,each:=range products{
    annualGrossProfit += each.Price*each.TotalSold*(1-each.Discount)
  }
  fmt.Printf("Annual Gross Profit: Rp %d\n",annualGrossProfit)

  // Data Type Exercise
  // ------------------

  // API Error Code
  // API := []map[int]string{
  //   // {code : description}
  //   {1 : "Incorrect Input"},
  //   {2 : "The server encounters internal error"},
  //   {4 : "The server is overloaded by too much traffic"},
  //   {8 : "You are not authorized to proceed with the input"},
  // }

  fmt.Print("Input: ")
  // var input int32
  // fmt.Scanln(&input)
  input := 3

  var output []string

  if input == 0{
    output = append(output, "No Error")
  }else{
    for input!=0{
      switch{
      case input>=8:
        output = append(output, "You are not authorized to proceed with the input")
        input -= 8
      case input>=4:
        output = append(output, "The server is overloaded by too much traffic")
        input -= 4
      case input>=2:
        output = append(output, "The server encounters internal error")
        input -= 2
      case input>=1:
        output = append(output, "Incorrect Input")
        input -= 1
      }
    }
  }
  fmt.Println("Output: ", output)

  // Closure and Scope
  callbacks := CreateCallbacks(5)
  for i:=0;i<5;i++{
    fmt.Println(callbacks[i]())
  }

  // English to Indonesia Dictionary CLI
  // dictionary := map[string]string{
  //   // ID:EN
  //   "membaca":"read",
  //   "minum":"drink",
  //   "tas punggung":"backpack",
  //   "tertawa":"laugh",
  //   "tidur":"sleep",
  // }

  // for{
  //   fmt.Println("ID to EN Dictionary")
  //   fmt.Println("Menu:")
  //   fmt.Println("1. Translate")
  //   fmt.Println("2. Add word")
  //   fmt.Println("3. Remove word")
  //   fmt.Println("4. Print dictionary")

  //   var input int64
  //   fmt.Print("Input: ")
  //   fmt.Scanln(&input)

  //   switch input{
  //   case 1:
  //     fmt.Print("Word to translate: ")
  //     var input string
  //     fmt.Scanln(&input)

  //     _, ok := dictionary[input]

  //     if ok{
  //       fmt.Println("ID: ",input)
  //       fmt.Println("EN: ", dictionary[input])
  //     }else{
  //       fmt.Printf("sorry, \"%s\" is not found in dictionary\n", input)
  //     }
  //   case 2:
  //     fmt.Print("Word to be added in dict: ")
  //     var input string
  //     fmt.Scanln(&input)

  //     // var id string
  //     // var en string
  //     // var isID bool = false

  //     // for _,char:=range input{
  //     //   if char == rune("#"){
  //     //     continue
  //     //   }
  //     //   if isID{
  //     //     id+=char
  //     //   }else{
  //     //     en+=char
  //     //   }
  //     // }
  //     result := strings.Split(input, "#")

  //     if len(result) == 1{
  //       fmt.Println("word must be separated with # char")
  //     }else{
  //       if _, ok := dictionary[result[0]]; ok{
  //         fmt.Printf("cannot add existing word \"%s\"\n", result[0])
  //       }else{
  //         dictionary[result[0]] = result[1]
  //         fmt.Println("new word successfully added")
  //       }
  //     }
  //   case 3:
  //     fmt.Print("Word to be removed: ")
  //     var input string
  //     fmt.Scanln(&input)
  //     if _,ok:=dictionary[input]; ok{
  //       delete(dictionary,input)
  //       fmt.Printf("\"%s\" has been removed\n", input)
  //     }else{
  //       fmt.Printf("sorry, \"%s\" is not found in dictionary\n", input)
  //     }
  //   case 4:
  //     var keys []string
  //     for k:=range dictionary{
  //       keys = append(keys, k)
  //     }
  //     sort.Strings(keys)
  //     for i,k:=range keys{
  //       fmt.Printf("%d. %s: %s\n", i+1, k, dictionary[k])
  //     }
  //   }
  //   fmt.Println()
  // }

  // Lot Billing
  var vehicle string = "Motorcycle"
  var duration int = 25
  var billing int

  // fmt.Print("Vehicle: ")
  // fmt.Scanln(&vehicle)
  // fmt.Print("Duration(in hour): ")
  // fmt.Scanln(&duration)

  if vehicle == "Motorcycle"{
    billing = 3000+(2000*(duration-1))
    if duration>24{
      billing += 20000
    }
  } else{ //Car
    billing = 7000+(5000*(duration-1))
    if duration>24{
      billing += 50000
    }
  }
  fmt.Printf("Billing: %d\n", billing)

  //POS - Point of Sales App
  // productList := []Products{
  //   {"00001", "Coca-cola", 3000},
  //   {"00002", "Sprite", 2500},
  //   {"00003", "Fanta", 2500},
  //   {"00004", "Instant Noodles", 3500},
  //   {"00005", "Coffee", 5000},
  // }

  // Robot Translator
  command := "RRAAALA"
  x := "abcd\n"
  y := "xyz\n"
  z := command+x+y
  fmt.Println(z)
  var result string
  var currentCommand rune
  var counter int = 1
  s := ""
  for _,c:=range command{
    if currentCommand == 0{
      currentCommand = c
    }else if currentCommand == c{
      counter++
    }else{ //currentCommand != c && currentCommand != 0
      s = ""
      if counter > 1{
        s = "s"
      }
      if currentCommand == []rune("R")[0]{
        result += fmt.Sprintf("Move right %d time%s\n", counter, s)
      }else if currentCommand == []rune("L")[0]{
        result += fmt.Sprintf("Move left %d time%s\n", counter, s)
      }else{ // c == []rune("A")[]
      result += fmt.Sprintf("Move advance %d time%s\n", counter, s)
      }
      counter = 1
      currentCommand = c
    }
    fmt.Println(c)
    fmt.Println(counter)
  }
  if currentCommand == []rune("R")[0]{
    result += fmt.Sprintf("Move right %d time%s\n", counter, s)
  }else if currentCommand == []rune("L")[0]{
    result += fmt.Sprintf("Move left %d time%s\n", counter, s)
  }else{ // c == []rune("A")[]
  result += fmt.Sprintf("Move advance %d time%s\n", counter, s)
  }
  fmt.Println(result)

  // Tic Tac Toe

  inputs := []rune("XOXXOOXXO")

  // var matrix [3][3]rune

  // for i:=0;i<3;i++{
  //   for j:=0;j<3;j++{
  //     matrix[i][j] = inputs[i+j]
  //   }
  // }

  // for i:=0;i<3;i++{
  //   fmt.Println(matrix[i])
  // }
  // state
  // X wins
  // O wins
  // Draw
  // Game in progress
  // Invalid grid

  // inputs = []rune("XOOXOXOXO")
  if  (inputs[0] == []rune("X")[0] && inputs[0] == inputs[1] && inputs[1] == inputs[2]) ||
      (inputs[3] == []rune("X")[0] && inputs[3] == inputs[4] && inputs[4] == inputs[5]) ||
      (inputs[6] == []rune("X")[0] && inputs[6] == inputs[7] && inputs[7] == inputs[8]) ||
      (inputs[0] == []rune("X")[0] && inputs[0] == inputs[3] && inputs[3] == inputs[6]) ||
      (inputs[1] == []rune("X")[0] && inputs[1] == inputs[4] && inputs[4] == inputs[7]) ||
      (inputs[2] == []rune("X")[0] && inputs[2] == inputs[5] && inputs[5] == inputs[8]) ||
      (inputs[0] == []rune("X")[0] && inputs[0] == inputs[4] && inputs[4] == inputs[8]) ||
      (inputs[2] == []rune("X")[0] && inputs[2] == inputs[4] && inputs[4] == inputs[6]){
        fmt.Println("X wins")
  } else if  (inputs[0] == []rune("O")[0] && inputs[0] == inputs[1] && inputs[1] == inputs[2]) ||
      (inputs[3] == []rune("O")[0] && inputs[3] == inputs[4] && inputs[4] == inputs[5]) ||
      (inputs[6] == []rune("O")[0] && inputs[6] == inputs[7] && inputs[7] == inputs[8]) ||
      (inputs[0] == []rune("O")[0] && inputs[0] == inputs[3] && inputs[3] == inputs[6]) ||
      (inputs[1] == []rune("O")[0] && inputs[1] == inputs[4] && inputs[4] == inputs[7]) ||
      (inputs[2] == []rune("O")[0] && inputs[2] == inputs[5] && inputs[5] == inputs[8]) ||
      (inputs[0] == []rune("O")[0] && inputs[0] == inputs[4] && inputs[4] == inputs[8]) ||
      (inputs[2] == []rune("O")[0] && inputs[2] == inputs[4] && inputs[4] == inputs[6]){
        fmt.Println("O wins")
  } else{
    fmt.Println("Draw")
  }
  // fmt.Println(inputs)
  // fmt.Println((inputs[0] == []rune("X")[0]))
  // fmt.Println(inputs[0] == inputs[3])
  // fmt.Println(inputs[3] == inputs[6])

  // ALGORITHMS DATA STRUCTURES
  // Some of most used algorithms: searching, sorting, recursive, hashing, bruteforce, etc
  // ALU -> Arithmetic Logic Unit
  // Bitwise operation usually used for low-level manipulation for data compression and encryption algorithms

  maxSum := -1 << 31 // it's -(1<<31), minimum value for int32
  fmt.Println(maxSum)

  // Classification of Data Structure
  // Linear Data Structure:
  // Non-Lin






  // Go Fundamentals
  // best practice to package names: short, lowercase, single-word. avoid underscores or mixedCaps
  // Ex: strconv, suffixarray,
  // It's better to create one package per directory except _test
  // module path format tipically <domain>/<project-description>

  // direct dependency imported directly in your code using import keyword
  // indirect dependency is required but you don't import it in your code, usually one of your direct dependency
  // has this module as requirement
  // to see the hierarcy of indirect dependencies on why they needed write this command
  // go mod why <indirect-dependecies>

  // go get example.com/module@v1.0.0 -> specific version
  // go get example.com/module@latest -> latest version

  // go list -m -u all -> check update of go module listed in go.mod
  // go mod tidy -> ensure go.mod file matches module used in your project. add missing modules and remove unused modules

  // error type in Go is implemented as
  // type Error interface {
    // Error() string
  // }

  // error usually returned as the last argument of a function/method

  a,b := 10,0
  r, err := divide(a,b)
  if err!=nil{
    switch{
    case errors.Is(err, ErrDivideByZero):
      fmt.Println("divide by zero error")
    default:
      fmt.Printf("unexpected error : %s\n", err)
    }
  }
  fmt.Println(r)

  // Wrapping error are mainly to make debugging easier(?), also can use errors.Is and errors.As
  // sometimes you don't need to wrapping an error since it's can lower security and privacy

  // error message shouldn't capitalized
  // errors.New variable should have format of errFoo or ErrFoo

  // unit test ideally should constitute bulk of tests
  // integration -> how piece of code functions when interact with other codes. these tests communicate with
  // external dependencies like database or external API

  // Given-When-Then is a part of Behavior-Driven Development(BDD)
  // FIRST (Fast, Isolated, Repeatable, Self-Validating, Timely)
  // bad unit test: manual assertion, multiple assertion, redundant assertion, coupling test (one unit test,
  // dependant on another unit test), will failing test, login in test (avoid any logic in test, since every logic
  // need it's own test, logic as if-else etc)
  // generally, avoid testing IO (e.g. files, database, network, request), Time, and Random number

  // OOP Concept
  // Golang doesn't have inheritance, instead it's using composition. Child can use parents method directly
  c1 := Marvel{
    Comic{
      "MCU",
    },
  }

  c2 := DC{
    Comic{
      "DC",
    },
  }

  fmt.Println(c1.ComicUniverse())
  fmt.Println(c2.ComicUniverse())

  // A child can have multiple parents btw

  // SOLID-> Single Responsibility principle(SRP), Open-Closed Principle(OCP), Liskov Subtitution Principle(LSP),
  // Interface Segregation Principle(ISP), Dependency Inversion Principle(DIP)

  // abstraction in Golang also implemented using interface
  // duck typing vs structural typing
  // duck typing (runtime) focus on what an object can do (used in Python, Ruby, JS)
  // structural typing (compile-time) focus on the structure (used in Go, Rust, TypeScript)
  // wait, Go use static typing with interface satisfaction (? what is this)

  // Oh so there's 2 typing: static and dynamic. duck is dynamic. Structural and nominal are static.

  // Go use static typing with interface satisfaction means if a type implements all method of an interface,
  // that type is satisfy that interface
  // example type A interface{x() y()}, type U struct{}, func U x(), func U y(), var a A := U{} -> valid

  // test double in golang

  // concurrency vs parallelism

  // Goroutine -> can be tricky to test or predict it's behavior
  // goroutine give you concurrency by design, and parallelism when possible

  i = 0
  printHitung := func(done chan bool){
    i++
    fmt.Print(i)
    done<-true
  }

  done := make(chan bool)

  for i:=65;i<75;i++{
    fmt.Print(string(rune(i)))
    go printHitung(done)
    <-done // wait until done receive true (done<-true)
  }

  // data race -> 2 goroutines, access same variable, at least one write, no synchronization

  // race condition -> mutex,

  var wg sync.WaitGroup
  wg.Add(2)

  go func(){
    // for range 100000{
    for i:=0; i<100000; i++{
      increment()
    }
    wg.Done()
  }()

  go func(){
    for i:=0; i<100000; i++{
      increment()
    }
    wg.Done()
  }()

  wg.Wait()
  fmt.Println("Counter with mutex: ",coun)

  // wg.add() tells how much goroutine we use. wg.Done basically decrement wg.add()--,
  // wg.Wait only continue when wg.Add = 0

  // fmt is safe for concurrent use but doesn't guarantee atomic or ordered output

  // best practice to always do close(channel)
  // and do it in the same scope when it's initialize(?)
}

var coun int
var mutex sync.Mutex
func increment(){
  mutex.Lock()
  defer mutex.Unlock()
  coun++
}
type Comic struct{
  Universe string
}

func (c Comic) ComicUniverse() string{
  return c.Universe
}

type Marvel struct{
  Comic
}
type DC struct{
  Comic
}

// you can define expected error
var ErrDivideByZero = errors.New("divide by zero")

func divide(a, b int) (int, error){
  if b==0{
    return 0, ErrDivideByZero
  }
  return a/b, nil
}

func DoSomething() error{
  return errors.New("create new error with static error message")
  // you can even use fmt.Errorf to add dynamic data type like int with %d etc
  // there's %w that wrap errorr within error, let's learn that one day
}

// type Products struct{
//   Barcode string
//   ProductName string
//   Price int
// }

// type PurchaseSummary struct{
//   Barcodes []string
//   Qty []int
//   SubTotal []int
//   Total int
// }

// func CalculatePurchase(barcodes []string) PurchaseSummary{

// }

// func GenerateReceiptText(ps PurchaseSummary){

// }

func CreateCallbacks(n int) (res []func()int){
  // for i:=0;i<n;i++{
  for i:=range n{
    res = append(res, func() int {return i})
  }
  return res
}

type variable struct{
  Product string
  Price int
  TotalSold int
  Discount int
}

type MyStruct struct{
  i int
}