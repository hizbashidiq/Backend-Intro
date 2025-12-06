package main

import "net/http"

const USERNAME = "batman"
const PASSWORD = "secret"

// in essence, middleware is a block of code that executed before or after http request
// in NodeJS and Rails it's called middleware, in Java Enterprise it's called filters, in C# it's called delegate handlers
// http.Handler interface is the most popular data type that used for middleware management

func Auth(w http.ResponseWriter, r *http.Request)bool{
  username,password,ok := r.BasicAuth()
  if !ok{
    w.Write([]byte("Something went wrong"))
    return false
  }

  isValid := (username == USERNAME) && (password == PASSWORD)
  if !isValid{
    w.Write([]byte("Wrong username/password"))
    return false
  }

  return true
}

func AllowOnlyGET(w http.ResponseWriter, r *http.Request) bool {
  if r.Method != "GET"{
    w.Write([]byte("Only GET request is allowed"))
    return false
  }
  return true
}

func MiddlewareAuth(next http.Handler) http.Handler{
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
    username, password, ok := r.BasicAuth()
    if !ok{
      w.Write([]byte("Something went wrong"))
      return
    }
    isValid := (username==USERNAME) && (password==PASSWORD)
    if !isValid{
      w.Write([]byte("Wrong username/password"))
      return
    }
    next.ServeHTTP(w, r)
  })
}

func MiddlewareAllowOnlyGET(next http.Handler) http.Handler{
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
    if r.Method != "GET"{
      w.Write([]byte("Only accept GET request"))
      return
    }
    next.ServeHTTP(w,r)
  })
}

type CustomMux struct{
  http.ServeMux
  middlewares []func(next http.Handler) http.Handler
}

func (c *CustomMux) RegisterMiddleware(next func(next http.Handler) http.Handler){
  c.middlewares = append(c.middlewares, next)
}

func (c *CustomMux) ServeHTTP(w http.ResponseWriter, r *http.Request){
  // every mux need to have ServeHTTP method
  // every 3rd party router in Go (Gin, Chi, Gorilla Mux, etc) using custom mux
  var current http.Handler = &c.ServeMux

  for _, next := range c.middlewares{
    current = next(current)
  }

  current.ServeHTTP(w, r)
}