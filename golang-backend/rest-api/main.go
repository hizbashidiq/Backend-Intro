package main

import (
	"encoding/json"
	"errors"
	"fmt"

	// "strings"

	// "net/http"
	"os"
	// "log"
	// "github.com/gin-gonic/gin"
	"context"
	"time"
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
  framework = "gin"
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

  // context -> easier to manage passing request-scoped value, cancellation signal, and deadline to all goroutine
  // WithValue(), WithCancel(), WithDeadline(), WithTimeout()
  // context interface method: Deadline(), Done(), Err(), Value()
  // best practices -> pass context as the first argument, always call CancelFunc to release the resource even
  // if it's succeed and not canceled (often using defer cancel())

  // background is the root of any context tree and never canceled

  // ctx := context.TODO()
  ctx := context.Background()

  ctx = context.WithValue(ctx, "myKey", "MyVal")

  // context.todo() and context.background() basically the same, it's just convey different intent of dev
  doSomething(ctx)

  // value stored in specific context is immutable

  // best practice-if you need a data to run a function, pass it as an argument instead store it in context

  // if more than one case fulfill the condition in select statement, it will be chosen randomly

  ctx = context.Background()
  doSomething1(ctx)

  // errors.New and fmt.Errorf basically the same but usually .New will be used


  for i:=1;i<4;i++{
    fmt.Printf("validating...%d", i)
    // if err:=validate(i);err == itsTwo{
    //   fmt.Println("this is error two")
    // }else if err!=nil{
    //   fmt.Println("there's an error")
    // }else{
    //   fmt.Println("valid!")
    // }
    err := runValidate(i)

    var valueErr *ValueError
    // if err == itsTwo || errors.Unwrap(err) == itsTwo || errors.Unwrap(errors.Unwrap(err)) == itsTwo{
    if errors.Is(err, itsTwo){
      // errors.Is using error type unwrap method which mean you still need to implement unwrap method in your custom
      // error struct
      fmt.Println("Oh no!")
    }else if errors.As(err, &valueErr){
      fmt.Printf("value error (%d): %v\n", valueErr.Value, valueErr.Err)
    }else if err!=nil{
      fmt.Println("there was an error :", err)
    }else{
      fmt.Println("valid!")
    }
  }

  // input validation should happen as early as possible in the data flow, preferably as soon as the data received
  // from the external party

  // syntatic validation should enforce correct syntax of structured fields (e.g. date, currency symbol)
  // semantic validation should enforce correct value regarding business context (e.g. start date before end date)

  // middleware usually used for logging, authentication/authorization, error handling, CORS, request throttling

  // if used properly, AES is effectively unbreakable in practice
  // when AES get hacked it never the AES, it mostly the human error; hardcodded key, keys committed to git,
  // business and product constraint, etc
  // security is often a negotiated compromise

  // most used cryptographic algorithm
  // SHA-256/SHA-512 -> data integrity, hashing non-password data
  // bcrypt/scrypt/Argon2 -> password hashing
  // AES-GCM -> encrypt database field
  // HMAC (usually HMAC-SHA256) -> message authentication
  // RSA, ECDSA, Ed25519 -> digital signatures (JWT, OAuth, auth systems)

  // AES, bcrypt/Argon2, HMAC, and TLS account for more than 90% of real world cryptography usage

  // encryption two way, hashing one way that's why hashing mostly used for password
  // you must select which one to use that suitable for your specific case

  // Authentication ->process of proving that some facts or documents are genuine
  // Authorization ->process of giving someone the ability to access a resource

  // JWT -> JSON Web Token
  // JWT part : header.payload.signature
  // header contains algorithm that used to sign the token, and it specifies type of token (in this case JWT)
  // payload carries the claim which are statements about the entity and additional data
  // 3 types of claim:
  // Registered claims
  // Public claims
  // Private claims
  // signature includes the result of a function that involves header, payload, and secret key (symmetric) or public
  // key and private key (asymm) used to sign the token. It's used to verify that the content of the token hasn't
  // been tampered

  // two parts of authentication using JWT with symmetric signing algorithm
  // signing a token, verifying a token

  // OAuth 2.0
  // 4 roles
  // The third-party application (app that require to access user information)
  // The resource server (API we want to access)
  // The Authorization server (might be the same server with resource server and might be different)
  // The resource owner (user)

  // different scenario will have different OAuth2.0 flow

  // ID token, access token, refresh token

  // Authenticator app is not OAuth 2.0.
  // refresh tokens can be refreshed indefinitely
  // refresh token must be stateful on the server so server can act on it depend on it's state

  // in go, string and []byte basically the same (slices of bytes) except the mutability. strings are read only

  // Authentication Factors
  // Something you know e.g. password, PIN
  // Something you have e.g. smartphone
  // Something you are e.g. biometric authentication
  // 2FA : something you know + something you have

  // IAM (Identity and Access Management)

  // jwt payload/claims are basically up to dev. exp is mandatory in practice. sub is strongly recommended


  // Clean architecture aim to separate software into layers
  // Dependency Rule : Dependency can only point inward. Nothing in inner circle can know anything about
  // outer circle. e.g. a var, functions, class name in outer circle must not be mentioned in inner circle.
  // Inner to Outer
  // 1. Model/Entity Layer (business rule): core business concepts, no framework, pure go construct and method
  // 2. Repository Layer
  // 3. Service/Usecase Layer (application rules): what the system does, define interface
  // 4. Handler/Controller Layer

  // repository: how do I get/save data? (DB handler essentially)
  // use case: what does the app do? (business logic)

  // interface in go: to decouple behavior from implementation when change is expected

  // Documentation purpose
  // 1. Reference and functionality
  // 2. Guides and tutorials
  // 3. Examples and use case

  // Stripe's API doc pretty much standard (?)

  // graceful shutdown in go http server:
  // listen to incoming OS signal such as SIGTERM and SIGINT
  // call .shutdown() to our HTTP server

  // if there's no in-flight request, the server will shutdown immediately
  // if there's in-flight request, the server will wait until that request is completed
  // if there's new request after initiate graceful shutdown, the new request will not be served
  // 
}


// custom error type usually has a suffix of Error
type ValueError struct{
  Value int
  Err error
}

func newValueError(value int, err error) *ValueError{
  return &ValueError{
    Value : value,
    Err : err,
  }
}
func (ve *ValueError) Error() string{
  return fmt.Sprintf("value error: %s", ve.Err)
}
func (ve *ValueError) Unwrap() error{
  return ve.Err
}

var itsTwo = errors.New("it's 2")

func runValidate(i int) error{
  if err := validate(i); err!=nil{
    return fmt.Errorf("run error: %w", err)
  }
  return nil
}

func validate(i int) error{
  if i == 1{
    // return fmt.Errorf("it's 1")
    return newValueError(i, fmt.Errorf("it's 1"))
  }else if i == 2{
    // return itsTwo
    return newValueError(i, itsTwo)
  }else{
    return nil
  }
}

func doSomething1(ctx context.Context){
  ctx, cancelCtx := context.WithCancel(ctx)

  printCh := make(chan int)

  go doAnother1(ctx, printCh)
  for i:=0;i<=3;i++{
    printCh<-i
  }
  cancelCtx()
  time.Sleep(100*time.Millisecond)
  fmt.Printf("doSomething1: finished\n")

}

func doAnother1(ctx context.Context, ch <-chan int){ //it's kinda a worker
  for{
    select{
    case <-ctx.Done():
      if err:=ctx.Err(); err!=nil{
        fmt.Printf("Do Another1: Error:%s\n", err)
      }
      fmt.Printf("Do Another1: Finished\n")
      return
    case num:=<-ch:
      fmt.Printf("doAnother1: %d\n", num)
    }
  }
}

func doSomething(ctx context.Context){
  fmt.Printf("Do Something: myKey's value is %s\n", ctx.Value("myKey"))

  anotherCtx := context.WithValue(ctx, "myKey", "anotherVal")
  doAnother(anotherCtx)
  // when using ctx.value, it will find outer most wrapped value for the given key

  fmt.Printf("Do Something: myKey's value is %s\n", ctx.Value("myKey"))
}

func doAnother(ctx context.Context){
  fmt.Printf("Do Another: myKey's value is %s\n", ctx.Value("myKey"))
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