package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

var mux = http.NewServeMux()

var sDefault = &http.Server{
  Addr: ":8080",
  Handler: mux,
}

func runDefault(){
  mux.HandleFunc("/hellodefault", func(w http.ResponseWriter, r *http.Request){
    w.WriteHeader(http.StatusOK)
    switch r.Method{
    case "GET":
      fmt.Fprintln(w, "WORLD")
    default:
      http.NotFound(w, r)
    }
  })

  mux.HandleFunc("/marketdefault/", func(w http.ResponseWriter, r *http.Request){
    itemID := strings.TrimPrefix(r.URL.Path, "/marketdefault/")
    switch r.Method{
    case "GET":
      w.WriteHeader(http.StatusOK)
      fmt.Fprintln(w, "You request this item:", itemID)
    case "PUT":
      w.WriteHeader(http.StatusAccepted)
      fmt.Fprintln(w, "You want to edit this item:", itemID)
    default:
      http.NotFound(w, r)
    }
  })
  log.Fatal(sDefault.ListenAndServe())
}