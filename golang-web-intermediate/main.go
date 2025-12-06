package main

import (
	// "golang-web-intermediate/conf"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"
  "log"
  "strings"
	gubrak "github.com/novalagung/gubrak/v2"
)

func main(){

  // var tmpl, err = template.ParseGlob("views/*") // this need to be in the beginning since it's parsing
  // // so next process will just render this

  // if err != nil {
  //   panic(err.Error())
  // }

  // http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
  //   var filepath = path.Join("views", "index.html")
  //   // path.Join usually used for non-filesystem path like URL and web
  //   // filepath.Join used for filesystem path because it's changed based on user operating system
  //   var tmpl, err = template.ParseFiles(filepath)
  //   if err != nil{
  //     http.Error(w, err.Error(), http.StatusInternalServerError)
  //     return
  //   }

  //   // parsing data to html can be done using map or struct
  //   // map -> key become variable name
  //   // struct -> property become variable name
  //   var data = map[string]interface{}{
  //     "title":"Golang Web",
  //     "name":"John",
  //   }

  //   err = tmpl.Execute(w, data)
  //   if err != nil{
  //     http.Error(w, err.Error(), http.StatusInternalServerError)
  //   }

  // }) //2nd parameter can also be closure or anonymous function
  http.HandleFunc("/", handlerIndex)
  http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request){
    var data = M{"name":"John Wick"}
    var tmpl = template.Must(template.ParseFiles(
      "views/index.html",
      "views/_header.html",
      "views/_message.html",
    ))
    err := tmpl.ExecuteTemplate(w, "index", data)
    // ExecuteTemplate used for ParseGlob since it's takes multiple template, executetemplate let we use specific template
    // 2nd parameter is template name (not file name)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
  })

  http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request){
    var data = M{"name":"John Wick"}

    var tmpl = template.Must(template.ParseFiles(
      "views/about.html",
      "views/_message.html",
      "views/_header.html",
    ))

    eer := tmpl.ExecuteTemplate(w, "about", data)
    if eer != nil {
      http.Error(w, eer.Error(), http.StatusInternalServerError)
    }
  })

  http.HandleFunc("/hello", handlerHello)
  http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request){
    var message = "DATA"
    w.Write([]byte(message))
  })
  http.Handle("/static/",
    http.StripPrefix("/static/", // it's strip /static/ from source so the path gonna be ./assets/ instead of
    // ./assets/static/ so static become purely route
      http.FileServer(http.Dir("assets")))) //http.Dir() is to adjust path parameter separator like in windows using \



  http.HandleFunc("/actionvariable", func(w http.ResponseWriter, r *http.Request){
    var data = Person{
      Name : "John Wick",
      Gender : "Male",
      Hobbies : []string{"get in trouble", "solo", "motorcycle"},
      Info : Info{"High Table", "Hotel Continental"},
    }

    var tmpl = template.Must(template.ParseFiles(path.Join("views","view.html")))
    if err := tmpl.Execute(w, data); err != nil{
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
  })

  http.HandleFunc("/function", func(w http.ResponseWriter, r *http.Request){
    var data = Superhero{
      Name: "Bruce Wayne",
      Alias: "Batman",
      Friends: []string{"Superman", "Flash", "Green Lantern"},
    }

    tmpl := template.Must(template.ParseFiles(path.Join("views", "function.html")))
    err := tmpl.Execute(w, data)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }

  })

  var funcMap = template.FuncMap{
    "unescape": func(s string) template.HTML{
      return template.HTML(s)
    },
    "avg": func (n ...int) float64{
      var sum int = 0
      // for i:=range n{ -> if only use 1 variable, it's return index
      //   sum +=
      // }
      for _, each:=range n{
        sum += each
      }
      return float64(sum)/float64(len(n))
    },
  }


  http.HandleFunc("/custfunc", func(w http.ResponseWriter, r *http.Request){
    tmpl := template.Must(template.New("custfunc.html"). //it's different with parsefiles, you don't need path
      Funcs(funcMap).
      // ParseFiles("custfunc.html"))
      ParseFiles(path.Join("views","custfunc.html")))
      // parsefiles here != with normal parsefiles
      // the other ParseFiles is a function from template package
      // while the this one is a method from *template.Template
      // so it's a method after we create template.Template instances using template.New()
    if err := tmpl.Execute(w, nil); err != nil{
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
  })

  http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request){
    tmpl := template.Must(template.New("home").ParseFiles(path.Join("views", "renderspec.html")))
    if err := tmpl.Execute(w, nil); err != nil{
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
  })

  http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request){
    tmpl := template.Must(template.New("test").ParseFiles(path.Join("views","renderspec.html")))
    if err:= tmpl.Execute(w, nil); err!=nil{
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
  })

  http.HandleFunc("/htmlstring", func(w http.ResponseWriter, r *http.Request){
    tmpl := template.Must(template.New("main-template").Parse(view))
    if err := tmpl.Execute(w, nil); err!=nil{
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
  })

  http.HandleFunc("/redirect", func(w http.ResponseWriter, r *http.Request){
    http.Redirect(w,r, "/htmlstring", http.StatusTemporaryRedirect)
  })

  http.HandleFunc("/httpmethod", func(w http.ResponseWriter, r *http.Request){
    switch r.Method{
    case "POST":
      w.Write([]byte("post"))
    case "GET":
      w.Write([]byte("get"))
    default:
      http.Error(w,"", http.StatusBadRequest)
    }
  })

  http.HandleFunc("/input", routeInput)
  http.HandleFunc("/result", routeResult)

  http.HandleFunc("/input1", routeInput1)
  http.HandleFunc("/result1", routeResult1)

  http.HandleFunc("/formjson", routeFormJson)
  http.HandleFunc("/savejson", routeSaveJson)
  // http.Handle("/static/",
  //   http.StripPrefix("/static",
  //     http.FileServer(http.Dir("assets"))))

  http.HandleFunc("/jsonresponse", routeJsonResponse)

  // MultipartReader -> uploaded file aren't store in temporary local file first, not like ParseMultipartForm
  // instead file is directly taken from io.reader stream
  // payload AJAX request -> data you send from the browser to the server
  // in GET request, payload usually in query string /user?id=123
  // in POST request, it can be form data, json, raw text, etc

  http.HandleFunc("/uploadmultiple", handleUploadMultiple)
  http.HandleFunc("/processmultiple", handleProcessMultiple)

  http.HandleFunc("/files", handleFiles)
  http.HandleFunc("/list-files", handleListFiles)
  http.HandleFunc("/download", handleDownload)

  // multiplexer (mux) is a router. Every routing must be done through mux object
  // http.HandleFunc is using default mux. basically it's http.DefaultServeMux.HandleFunc()

  mux := http.DefaultServeMux
  // http.HandleFunc("/student", actionStudent)
  mux.HandleFunc("/student", actionStudent)

  var handler http.Handler = mux
  handler = MiddlewareAuth(handler)
  handler = MiddlewareAllowOnlyGET(handler)

  muxCustom := new(CustomMux)
  muxCustom.HandleFunc("/studentc", actionStudent)
  muxCustom.RegisterMiddleware(MiddlewareAuth)
  muxCustom.RegisterMiddleware(MiddlewareAllowOnlyGET)

  // router := new(ConfMux)
  // router.HandleFunc("/conf", func(w http.ResponseWriter, r *http.Request){
  //   w.Write([]byte("Hello"))
  // })
  // router.HandleFunc("/how", func(w http.ResponseWriter, r *http.Request){
  //   w.Write([]byte("How are you?"))
  // })

  http.HandleFunc("/cookie", actionCookie)
  http.HandleFunc("/delete", actionDelete)

  http.HandleFunc("/cancel", handleCancel)

  var address = "localhost:9000" //or 0.0.0.0:9000 or :9000
  fmt.Println("Server started at ", address)


  // AJAX -> Asynchronous JavaScript and XML
  // the differences between Form Data and Payload JSON is in the header Content-Type and structure of information
  // Form Data VS Payload JSON
  // application/x-www-form-urlencoded | application/json

  // first method to start a server
  err := http.ListenAndServe(address, nil) //2nd parameter is mux/multiplexer object
  // err := http.ListenAndServe(address, muxCustom)
  // ListenAndServe() is blocking so the code under this will not executed unless it stop or error




  // if err!=nil{
  //   fmt.Println(err)
  // }

  // second method to start a server
  var server = new(http.Server)
  server.Addr = address

  server.Handler = handler
  // server.Handler = muxCustom


  // server.Handler = router
  // server.ReadTimeout = conf.Configuration().Server.ReadTimeout*time.Second
  // server.WriteTimeout = conf.Configuration().Server.WriteTimeout*time.Second
  // server.Addr = fmt.Sprintf("localhost:%d", conf.Configuration().Server.Port)

  // if conf.Configuration().Log.Verbose{
  //   log.Printf("Starting server at %s \n", server.Addr)
  // }

  // err := server.ListenAndServe()
  if err!=nil{
    fmt.Println(err)
  }

  // routing in Go can be done by
  // 1. using http.HandleFunc()
  // 2. implement http.Handler interface in a struct then used in http.Handle()
  // 3. Create new multiplexer using http.ServeMux struct


}

func handleCancel(w http.ResponseWriter, r *http.Request){
  done := make(chan bool)
  go func(){
    time.Sleep(10*time.Second)

    done<-true
  }()

  select{
  case <-r.Context().Done():
    if err := r.Context().Err(); err!=nil{
      if strings.Contains(strings.ToLower(err.Error()), "canceled"){
        log.Println("Request cancelled")
      } else{
        log.Println("Unknown error occured.", err.Error())
      }
    }
  case <-done:
    log.Println("DONE")
  }
}

// type ConfMux struct{
//   http.ServeMux
// }

// func (c ConfMux) ServeHTTP(w http.ResponseWriter, r *http.Request){
//   if conf.Configuration().Log.Verbose{
//     log.Println("Incoming request from", r.Host, "accessing", r.URL.String())
//   }
//   c.ServeMux.ServeHTTP(w,r)
// }

var cookieName = "CookieData"

func actionDelete(w http.ResponseWriter, r *http.Request){
  c := &http.Cookie{}
  c.Name = cookieName
  c.Expires = time.Unix(0,0) //duration in time.Time
  c.MaxAge = -1 //duration in seconds
  http.SetCookie(w, c)

  http.Redirect(w,r,"/cookie", http.StatusTemporaryRedirect)
}

func actionCookie(w http.ResponseWriter, r *http.Request){
  cookieName := "CookieData"

  c := &http.Cookie{}

  if storedCookie,_ := r.Cookie(cookieName); storedCookie!=nil{
    c = storedCookie
  }

  if c.Value == ""{
    c = &http.Cookie{}
    c.Name = cookieName
    c.Value = gubrak.RandomString(32)
    c.Expires = time.Now().Add(5 * time.Minute)
    http.SetCookie(w, c)
  }

  w.Write([]byte(c.Value))
}

func actionStudent(w http.ResponseWriter, r *http.Request){
  // if !Auth(w, r){
  //   return
  // }
  // if !AllowOnlyGET(w,r){
  //   return
  // }

  if id := r.URL.Query().Get("id"); id!=""{
    OutputJSON(w, SelectStudent(id))
    return
  }

  OutputJSON(w, GetStudents())
}

func OutputJSON(w http.ResponseWriter, o interface{}){
  res, err := json.Marshal(o)
  if err!=nil{
    w.Write([]byte(err.Error()))
    return
  }

  w.Header().Set("Content-Type","application/json")
  w.Write(res)
}

func handleDownload(w http.ResponseWriter, r *http.Request){
  // recommended to always use return even if it's optional in handlefunc error handling
  // download implementation are basically the same across all programming language
  // by dealing with Content-Disposition header in HTTP response
  if err := r.ParseForm(); err!=nil{
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  path := r.FormValue("path")
  f, err := os.Open(path)
  if f!=nil{
    defer f.Close()
  }

  if err!=nil{
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  contentDisposition := fmt.Sprintf("attachment; filename=%s", f.Name())
  w.Header().Set("Content-Disposition", contentDisposition)

  if _,err := io.Copy(w, f); err!=nil{
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

}

func handleListFiles(w http.ResponseWriter, r *http.Request){
  files := []M{}
  basePath, _ := os.Getwd()
  filesLocation := filepath.Join(basePath, "files")

  err := filepath.Walk(filesLocation, func(path string, info os.FileInfo, err error) error{
    if err!=nil{
      return err
    }

    if info.IsDir(){
      return nil
    }

    files = append(files, M{"filename":info.Name(), "path":path})
    return nil
  })

  if err!=nil{
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  res, err := json.Marshal(files)
  if err!=nil{
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(res)
}

func handleFiles(w http.ResponseWriter, r *http.Request){
  tmpl := template.Must(template.ParseFiles(path.Join("views","download.html")))
  if err:=tmpl.Execute(w,nil); err!=nil{
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func handleProcessMultiple(w http.ResponseWriter, r *http.Request){
  if r.Method != "POST"{
    http.Error(w, "Only accept POST request", http.StatusBadRequest)
    return
  }

  basePath, _ := os.Getwd()
  reader, err := r.MultipartReader()
  if err!=nil{
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  for{
    part, err := reader.NextPart()
    if err == io.EOF{
      break
    }

    fileLocation := filepath.Join(basePath, "files", part.FileName())
    dst, err := os.Create(fileLocation)
    if dst!=nil{
      defer dst.Close()
    }
    if err!=nil{
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }

    if _, err:= io.Copy(dst, part); err!=nil{
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
  }

  w.Write([]byte("All files uploaded!"))
}

func handleUploadMultiple(w http.ResponseWriter, r *http.Request){
  // tmpl := template.Must(template.New("multifile.html").ParseFiles(path.Join("views","multifile.html")))
  tmpl := template.Must(template.ParseFiles(path.Join("views","multifile.html")))
  // use template.New if you need better control, setting custom functions (.Funcs), etc
  // if you simply load html file just use normal ParseFiles, most of the time you gonna use this

  if err := tmpl.Execute(w, nil); err!=nil{
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func routeJsonResponse(w http.ResponseWriter, r *http.Request){
  data := [] struct{
    Name string
    Age int
  }{
    {"A B", 20},
    {"B C", 21},
    {"C D", 22},
    {"D E", 23},
  }

  jsonInBytes, err := json.Marshal(data)
  if err!=nil{
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(jsonInBytes)

  // err := json.NewEncoder(w).Encode(data)
  // if err != nil {
  //     http.Error(w, err.Error(), http.StatusInternalServerError)
  //     return
  // }


}

func routeSaveJson(w http.ResponseWriter, r *http.Request){
  if r.Method == "POST"{
    decoder := json.NewDecoder(r.Body)
    payload := struct{
      Name string `json:"name"`
      Age int `json:"age"`
      Gender string `jsong:"gender"`
    }{}
    if err := decoder.Decode(&payload); err!=nil{
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }

    message := fmt.Sprintf("Hello, my name is %s. I'm years old %s.",
                            payload.Name,
                            // payload.Age,
                            payload.Gender)

    w.Write([]byte(message))
    return
  }
  http.Error(w,"Only accept POST request",http.StatusBadRequest)
}

func routeFormJson(w http.ResponseWriter, r *http.Request){
  // tmpl := template.Must(template.New("payload.html").ParseFiles(path.Join("views", "payload.html")))
  tmpl := template.Must(template.ParseFiles(path.Join("views","payload.html")))
  if err := tmpl.Execute(w, nil); err!=nil{
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }

}

func routeResult1(w http.ResponseWriter, r *http.Request){
  if r.Method != "POST"{
    http.Error(w,"", http.StatusBadRequest)
    return
  }

  // ParseMultipartForm is to parsing form data that been sent
  // 1024 is max memory, if the file exceeding that, it'll be saved in temporary file
  if err:=r.ParseMultipartForm(1024);err!=nil{
    http.Error(w,err.Error(), http.StatusInternalServerError)
    return
  }

  alias := r.FormValue("alias")
  uploadedFile, handler, err := r.FormFile("file")
  if err!=nil{
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  defer uploadedFile.Close()

  dir, err := os.Getwd()
  if err!=nil{
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  filename := handler.Filename
  if alias != ""{
    filename = fmt.Sprintf("%s%s", alias, filepath.Ext(handler.Filename))
  } //filepath.Ext to get file extension

  fileLocation := filepath.Join(dir,"files",filename)
  targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
  // WRONLY -> write only
  if err!=nil{
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  defer targetFile.Close()

  if _,err := io.Copy(targetFile, uploadedFile); err != nil{
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Write([]byte("DONE"))
}

func routeInput1(w http.ResponseWriter, r *http.Request){
  if r.Method != "GET"{
    http.Error(w,"",http.StatusBadRequest)
    return
  }

  tmpl := template.Must(template.New("upload.html").ParseFiles(path.Join("views","upload.html")))
  if err := tmpl.Execute(w, nil); err!=nil{
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func routeInput(w http.ResponseWriter, r *http.Request){
  if r.Method == "GET"{
    tmpl := template.Must(template.New("form").ParseFiles(path.Join("views","form.html")))
    if err:= tmpl.Execute(w, nil); err!=nil{
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    return
  }
  http.Error(w, "", http.StatusBadRequest)
}

func routeResult(w http.ResponseWriter, r *http.Request){
  if r.Method == "POST"{
    tmp := template.Must(template.New("result").ParseFiles(path.Join("views", "form.html")))
    if err := r.ParseForm(); err!=nil{
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }

    var name = r.FormValue("name")
    var message = r.Form.Get("message")

    // r.FormValue() and r.Form.Get() are basically the same here
    // in details, r.FormValue() are more simple while r.Form.Get() more advance and offer strict parsing control
    // most of the time just use r.FormValue() especially at early phase of learning (?)
    var data = map[string]string{"name":name, "message":message}

    if err := tmp.Execute(w, data); err!=nil{
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    return
  }

  http.Error(w, "", http.StatusBadRequest)
}

const view string = `<html>
  <head>
    <title>Template</title>
  </head>
  <body>
    <h1>Hello</h1>
  </body>
</html>`

type Superhero struct{
  Name string
  Alias string
  Friends []string
}

func (s Superhero) SayHello(from string, message string) string{
  return fmt.Sprintf("%s said : \"%s\"", from, message)
}

func (t Info) GetAffiliationDetailInfo() string{
  return "Winston Scott as Continental Manager"
}

type Info struct{
  Affiliation string
  Address string
}

type Person struct{
  Name string
  Gender string
  Hobbies []string
  Info Info
}

type M map[string]interface{}

func handlerIndex(w http.ResponseWriter, r *http.Request){
	var message = "welcome"
	w.Write([]byte(message))
}

func handlerHello(w http.ResponseWriter, r *http.Request){
	var message = "HEWWO"
	w.Write([]byte(message))
}