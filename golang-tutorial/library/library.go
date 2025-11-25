package library

import "fmt"

// func SayHello(name string){  //public
//   fmt.Println("Hello")
//   introduce(name)
// }

// func introduce(name string ){ //private
//   fmt.Println("Nama Saya ", name)
// }

// type student struct{
type Student struct{
  Name string
  // grade int
  Grade int
}

var orang = struct{
  Name string
  Age int
}{}

func init(){ //special function, this will executed before main()
  orang.Name = "John Wick"
  orang.Age = 22

  fmt.Println(orang)
}