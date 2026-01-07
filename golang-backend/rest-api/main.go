package main

import (
	"encoding/json"
	"fmt"
	// "strings"

	// "net/http"
	"os"
	// "log"
	// "github.com/gin-gonic/gin"
)

func main(){
	// url must avoid action. use method to categorized action
  // use kebab-case for path segments (not necessarily the name of path parameter)
  // use snake_case for query parameter
  // for error response (4xx and 5xx) provide error details in the response body i.e. error message, error code
  // links to resources for troubleshooting, etc
  // avoid overuse of 200 OK. use 200 OK for successful GET requests. For other successful operations, use status
  // code that accurately describe the action (201 created, 204 No content, 202 accepted etc)
  // Avoid Custom Status Code unless really necessary

  u := User{
    Name: "Erick",
    Age: 0,
    IsActive: true,
  }
  jsonBytes, err := json.Marshal(u)
  if err!=nil{
    fmt.Println(err)
  }
  fmt.Println(string(jsonBytes))

  enc := json.NewEncoder(os.Stdout)
  enc.Encode(u)
  fmt.Println(u)
  fmt.Printf("%T\n", u)

  emptySlice := []string{}
  var nilSlice []string
  jsonBytes, err = json.Marshal(nilSlice)
  if err!=nil{
    fmt.Println(err)
  }
  fmt.Println(string(jsonBytes))

  jsonBytes, err = json.Marshal(emptySlice)
  if err!=nil{
    fmt.Println(err)
  }
  fmt.Println(string(jsonBytes))

  // we can append nil slice but can't append nil map

  var usrObj User

  jsonStr := `{"Name":"Erick","Age":25,"IsActive":true}`
  if err = json.Unmarshal([]byte(jsonStr), &usrObj); err!=nil{
    fmt.Println(err)
  }
  fmt.Println(usrObj)

  // remember obj <-> byte <-> json when encoding or decoding using marshal or unmarshal method

  // struct tag mainly used to provide information to external tools or library, also to specify in how it
  // properties name as a json object, providing validation rules (?), and indicating the order of fields
  // when marshalling or unmarshalling

  // in golang you don't really need http framework, since the default already powerful
  // but to fasten development
  // some: Gin, echo, Beego, Fiber, Gorilla, etc

  // since default net/http is already powerful, for my first project just use net/http first
  // so you can learn how and why something works. after you comfortable with net/http then you can
  // refactor it into Gin or Chi (Gin is a lot faster but a lot of magic, while Chi are more idiomatic Go)
  // (Gin even considered fastest framework because of httprouter)
  // Chi -> microservice, full control
  // Gin -> big app, REST API that need high level performance and complete feature
  framework := ""
  // framework = "default"
  // framework = "gin"
  if framework == "default"{
    runDefault()
  }
  if framework == "gin"{
    runGin()
  }

  // ENV -> environment variables
  value := envVariable("APP_ENV")

  fmt.Printf("os package: name = %s \n", value)
  fmt.Printf("environment = %s \n", os.Getenv("APP_ENV"))

  // mostly you'll only need to read ENV, you'll not need to manage it's lifecycle
}

func envVariable(key string) string{

  os.Setenv(key, "production")
  return os.Getenv(key)
}

type User struct{
  Name string   `json:"first_name"`
  Age int       `json:"age,omitempty"` //omit value from json encoding if it's zero value (0 for int, "" for str, etc)
  // age,omitempty -> should be no space at all
  IsActive bool `json:"-"` //will be ignored when encode to json
}

type Account struct{
  Username string `json:"username" binding:"required"`
  Password string `json:"password" binding:"required"`
}