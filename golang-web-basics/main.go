package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"

	// "html/template"
	"database/sql" // this is for any sql be it mysql, postgresql, mariaDB, etc just change sql.Open parameter
	"encoding/json"
	"net/url"

	// sql.Open(driverName, connectionString)
	// example using postgresql:
	// sql.Open("postgres", "user=postgres password=secret dbname=test sslmode=disable")
	_ "github.com/go-sql-driver/mysql"

	// go get go.mongodb.org/mongo-driver/mongo

	"context"
	"log"

	// "time"
	"math"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	// go get github.com/strechr/testify
	// "runtime"
	"sync"
	// "time"
  "path/filepath"
  "crypto/md5"
)

func main(){

  // http.HandleFunc used for routing, yep i see

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
	// 	fmt.Fprintln(w, "HEWWO TOO")
	// })

	// http.HandleFunc("/index", index)

  // http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
  //   var data = map[string]string{
  //     "Name": "John Wick",
  //     "Message": "Man of focus, commitment, and sheer will. Something you know very little about",
  //   }

  //   var t, err = template.ParseFiles("template.html")
  //   if err != nil{
  //     fmt.Println(err)
  //     return
  //   }

  //   t.Execute(w, data)
  // })


  var urlString = "http://kalipare.com:80/hello?name=john wick&age=27"
  var u, e = url.Parse(urlString)

  if e != nil{
    fmt.Println(e)
    return
  }

  fmt.Printf("url: %s\n", urlString)
  fmt.Printf("protocol: %s\n", u.Scheme)
  fmt.Printf("host: %s\n", u.Host)
  fmt.Printf("path: %s\n", u.Path)

  var name = u.Query()["name"][0]
  var age = u.Query()["age"][0]
  fmt.Println(name, age)

  var jsonString = `{"Name": "john wick", "Age": 27}`
  // `` (backtick) used for something that literal all string even the escape character etc
  var jsonData = []byte(jsonString)

  var data User

  e = json.Unmarshal(jsonData, &data) // only accept []byte
  if e != nil {
    fmt.Println(e)
    return
  }

  fmt.Println(data.Fullname, data.Age)

  var data1 map[string]interface{}
  json.Unmarshal(jsonData, &data1)
  fmt.Println(data1["Name"], data1["Age"])

  var data2 interface{}
  json.Unmarshal(jsonData, &data2)

  var ddata2 = data2.(map[string]interface{}) // type assertion
  fmt.Println(ddata2["Name"], ddata2["Age"])

  jsonString = `[
    {"Name": "John Wick", "Age":27},
    {"Name": "Ethan Hunt", "Age":28}
  ]`
  jsonData = []byte(jsonString)

  var data3 []User

  json.Unmarshal(jsonData, &data3)
  fmt.Println(data3)
  fmt.Println(data3[0].Fullname)

  var object = []User{{"John Wick", 30}, {"Ethan Hunt", 40}}

  jsonData, _ = json.Marshal(object)
  jsonString = string(jsonData)
  fmt.Println(jsonString)



  http.HandleFunc("/users", users)
  http.HandleFunc("/user", user)

  fmt.Println("starting web server at http://localhost:8080/")
	// http.ListenAndServe("127.0.0.1:8080", nil)

  // go get github.com/go-sql-driver/mysql

  sqlQuery()
  sqlQueryRow()
  sqlPrepare()

  // insert, update, delete using Exec()
  // sqlExec()

  // mongodb
  // insert()
  find()
  // update()
  // find()
  // delete()
  aggregate()

  // in Go, testing file should named {file_name}_test.go
  // go test main.go main_test.go -v
  // -v for verbose
  // add "-bench=." to check benchmark
  // BenchmarkLuas-4         1000000000               0.3832 ns/op
  // means the function got run 1 billion times. in average it takes 0.3832 nano second to run 1 function

  var wg sync.WaitGroup
  for i:=0;i<5;i++{
    var data = fmt.Sprintf("data %d", i)

    wg.Add(1)
    go DoPrint(&wg, data)
  }
  wg.Wait() //blocking. Will wait all goroutines to finish before executing next code block
  // sync.WaitGroup is thread safe and will not have race condition
  // sync.WaitGroup used to specifically manage goroutines, not like channel
  // also performance wise sync.WaitGroup are better than channel
  // remember channel main function is to communicate or sharing data between goroutines

  runtime.GOMAXPROCS(2)

  // var wg sync.WaitGroup
  var meter Counter

  for i:=0;i<1000;i++{
    wg.Add(1)
    go func(){
      for j:=0;j<1000;j++{
        meter.Add()
      }
      wg.Done()
    }()
  }
  wg.Wait()
  fmt.Println(meter.Value())

  // to check race condition, add argument -race
  // ex: go run -race main.go

  // go mod vendor
  // go run -mod=vendor main.go
  // go build -mod=vendor -o executable
  // just make it easier and so we don't need to download dependency since it's saved in local
  // not a must, depends on what you need

  // CONCURRENCY PATTERN: PIPELINE
  // // NON-PIPELINE
  // // dummy file generator
  // log.Println("Start")
  // start := time.Now()

  // generateFiles()

  // duration := time.Since(start)
  // log.Println("done in", duration.Seconds(), "seconds")

  // // find md5 sum, rename file
  // start = time.Now()

  // proceed()

  // duration = time.Since(start)
  // log.Println("done in", duration.Seconds(), "seconds")

  // // PIPELINE
  // log.Println("Start")
  // start := time.Now()
  // // pipeline 1: loop all files and read it
  // chanFileContent := readFiles()

  // // pipeline 2: calculate md5sum
  // chanFileSum1 := getSum(chanFileContent)
  // chanFileSum2 := getSum(chanFileContent)
  // chanFileSum3 := getSum(chanFileContent)
  // chanFileSum := mergeChanFileInfo(chanFileSum1, chanFileSum2, chanFileSum3)

  // // pipeline 3: rename files
  // chanRename1 := rename(chanFileSum)
  // chanRename2 := rename(chanFileSum)
  // chanRename3 := rename(chanFileSum)
  // chanRename4 := rename(chanFileSum)
  // chanRename := mergeChanFileInfo(chanRename1, chanRename2, chanRename3, chanRename4)

  // // print output
  // counterRenamed := 0
  // counterTotal := 0

  // for fileInfo := range chanRename{
  //   if fileInfo.isRenamed{
  //     counterRenamed++
  //   }
  //   counterTotal++
  // }

  // log.Printf("%d/%d files renamed", counterRenamed, counterTotal)

  // duration := time.Since(start)
  // log.Println("Done in", duration.Seconds(), "seconds")

  // Okay so pipeline concurrency pattern need to do a lot of testing and trying, even if operation-wise is good
  // doesn't mean in practical it's good, you need to also consider I/O constraint, hardware specification, etc


  // Go Generics
  // alternative to go generics is using interface() with reflection API (?)

  total1 := Sum([]int{1,2,3,5}) // or Sum[int]([]int{1,2,3,5})
  fmt.Println(total1)

  total2 := Sum([]float32{1.5,2.5,3.5,4.5})
  fmt.Println(total2)

  total3 := Sum([]float64{1.5,2.5,3.5,4.5})
  fmt.Println(total3)

  // keyword "comparable" used to create generics function that compatible with all data type
  ints := map[string]int64{"first":12,"second":25}
  floats := map[string]float64{"first":12.5,"second":27.2}

  fmt.Println(Sum2(ints))
  fmt.Println(Sum2(floats))


}
// generic also can be used in struct
// generic can be used in function but can't be used in method
type UserModel[T int | float64] struct{
  Name string
  Scores []T
}

type Number interface{ //-> generic type constraint, can only be used in generic function
  // i.e. you can't declare Number as datatype in main()
  float64 | int64
}

func Sum3[K comparable, V Number](m map[K]V) V{
  var s V
  for _,x:=range m{
    s+=x
  }
  return s
}

func Sum2[K comparable, V int64 | float64](m map[K]V) V{
  var s V
  for _,x := range m{
    s+=x
  }
  return s
}

func Sum1(m map[string]int64) int64{
  var s int64
  for _,x := range m{
    s+=x
  }
  return s
}

func Sum[V int | float32 | float64](numbers []V) V{
  var total V
  for _,number:=range numbers{
    total+=number
  }
  return total
}

func rename(chanIn <-chan FileInfo) <-chan FileInfo{
  // every fan-out function essentially the same except the business logic or operation logic
  chanOut := make(chan FileInfo)

  go func(){
    for fileInfo := range chanIn{
      newPath := filepath.Join(TempPath, fmt.Sprintf("file-%s.txt", fileInfo.Sum))
      err := os.Rename(fileInfo.FilePath, newPath)
      fileInfo.IsRenamed = err==nil
      chanOut <- fileInfo
    }
    close(chanOut)
  }()
  return chanOut
}

func mergeChanFileInfo(chanInMany ... <-chan FileInfo)<-chan FileInfo{
  wg := new(sync.WaitGroup)
  chanOut := make(chan FileInfo)

  wg.Add(len(chanInMany))
  for _, eachChan := range chanInMany{
    go func(eachChan <-chan FileInfo){
      for eachChanData := range eachChan{
        chanOut <- eachChanData
      }
      wg.Done()
    }(eachChan)
  }

  go func(){
    wg.Wait()
    close(chanOut)
  }()
  return chanOut
}

func getSum(chanIn <-chan FileInfo) <-chan FileInfo{
  chanOut := make(chan FileInfo)

  go func(){
    for fileInfo := range chanIn{
      fileInfo.Sum = fmt.Sprintf("%x", md5.Sum(fileInfo.Content))
      chanOut <- fileInfo
    }
    close(chanOut)
  }()
  return chanOut
}

func readFiles() <-chan FileInfo{
  chanOut := make(chan FileInfo)

  go func(){
    err := filepath.Walk(TempPath, func(path string, info os.FileInfo, err error) error{
      if err != nil {
        return err
      }

      if info.IsDir(){
        return nil
      }

      buf, err := os.ReadFile(path)
      if err != nil {
        return err
      }

      chanOut <- FileInfo{
        FilePath: path,
        Content: buf,
      }
      return nil
    })
    if err != nil {
      log.Println("ERROR: ", err)
    }
    close(chanOut)
  }()
  return chanOut
}

type Counter struct{
  sync.Mutex
  Val int
}
func (c *Counter) Add(){
  // recommended to set Lock() and Unlock() in the object instead of main()
  c.Lock()
  c.Val++
  c.Unlock()
}
func (c *Counter) Value() int{
  return c.Val
}

func DoPrint(wg *sync.WaitGroup, message string){
  defer wg.Done()
  fmt.Println(message)
}

type Cube struct{
  Sisi float64
}
func (c Cube) Volume() float64{
  return math.Pow(c.Sisi,3)
}
func (c Cube) Luas() float64{
  return c.Sisi*c.Sisi*6
}
func (c Cube) Keliling() float64{
  return c.Sisi*12
}

func aggregate(){
  pipeline := make([]bson.M, 0)
  // it's messy to write pipeline in Go bson.D, so dev usually write it in json and convert it to bson

  // some of mongoDB aggregation
  // mongoDB pipelines each stage takes the previous results and transform it
  // $match -> like SQL where
  // $project -> like SQL select i.e. choose which field to output

  err := bson.UnmarshalExtJSON([]byte(strings.TrimSpace(`
    [
      {"$group":{
        "_id": null,
        "Total": {"$sum":1}
      }},
      {"$project":{
        "Total":1,
        "_id":0
      }}
    ]`)), true, &pipeline)
  if err != nil{
    log.Fatal(err)
  }

  db, err := connectMongo()
  if err != nil {
    log.Fatal(err)
  }

  csr, err := db.Collection("Student").Aggregate(ctx, pipeline)
  if err != nil{
    log.Fatal(err)
  }

  defer csr.Close(ctx)

  result := make([]bson.M, 0)
  for csr.Next(ctx){
    var row bson.M
    err := csr.Decode((&row))
    if err != nil {
      log.Fatal(err)
    }

    result = append(result, row)
  }
  fmt.Println(result)
}

func delete(){
  db, err := connectMongo()
  if err != nil{
    log.Fatal()
  }

  var filter = bson.M{"name":"King"}
  _, err = db.Collection("Student").DeleteOne(ctx, filter)
  if err != nil{
    log.Fatal(err)
  }
  fmt.Println("Delete success!")
}

func update(){
  db, err := connectMongo()
  if err != nil{
    log.Fatal(err)
  }

  var filter = bson.M{"name":"Wick"}
  var changes = Student{"JW01", "John Wick", 2}
  _, err = db.Collection("Student").UpdateOne(ctx, filter, bson.M{"$set":changes})
  if err != nil{
    log.Fatal(err)
  }
  fmt.Println("Update Success!")
}

func find(){
  db, err := connectMongo()
  if err != nil {
    log.Fatal(err)
  }

  // csr is cursor
  csr, err := db.Collection("Student").Find(ctx, bson.M{"name":"Wick"})
  // bson.M is alias for map[string]interface{}
  if err != nil {
    log.Fatal(err)
  }

  defer csr.Close(ctx)

  result := make([]Student, 0)
  for csr.Next(ctx){
    var row Student
    err := csr.Decode((&row))
    if err != nil {
      log.Fatal(err)
    }

    result = append(result, row)
  }
  if len(result)>0{
    fmt.Println(result)
  }
}

func insert(){
  db, err := connectMongo()
  if err != nil{
    log.Fatal((err))
  }
  // mongo-go-driver
  _, err = db.Collection("Student").InsertOne(ctx, Student{"W001", "Wick", 3})
  if err != nil{
    log.Fatal(err)
  }

  _, err = db.Collection("Student").InsertOne(ctx, Student{"K001", "King", 2})
  if err != nil{
    log.Fatal(err)
  }

  // there's no close connection in mongoDB?
  fmt.Println("Insert success!")

}

var ctx = context.Background()

func connectMongo() (*mongo.Database, error){
  clientOptions := options.Client()
  clientOptions.ApplyURI("mongodb://localhost:27017")
  client, err := mongo.NewClient(clientOptions)
  if err != nil {
    return nil, err
  }
  err = client.Connect((ctx))
  if err != nil {
    return nil, err
  }

  return client.Database("golang_tutorial"), nil

}

func sqlExec(){
  db, err := connect()
  if err != nil {
    fmt.Println(err)
    return
  }
  defer db.Close()

  // it's okay to use lowercase for sql

  // using exec only
  _, err = db.Exec("insert into tb_student values (?, ?, ?, ?)", "S001", "Saitama", 24, 5)
  if err != nil {
    fmt.Println(err)
    return
  }
  // yes you need to write all that e != nil thing. it's normal. it's really is how go programmer write their code
  // also use err instead of e. Go convention of error is err.

  // some rule of thumb on should you print or just return err
  // return err : inside helper functions, inside services/logic, http handlers (via http.Error etc)
  // print err : main() or executables, background worker
  fmt.Println("insert success!")

  _, err = db.Exec("update tb_student set age = ? where id = ?", 25, "S001")
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println("update success!")

  // using prepare
  stmt, err := db.Prepare("insert into tb_student values (?, ?, ?, ?)")
  if err != nil {
    fmt.Println(err)
    return
  }

  stmt.Exec("K001", "King", 35, 4)
  fmt.Println("insert success!")
}

func sqlPrepare(){
  // db.Prepare is used for performance, safety (not so much in Go), and reuse
  // if you only runs the query once, don't bother
  // it's usually used to batch insert/update (or even delete) since you only need parse+compile sql once
  // select doesn't really matter but this one just a tutorial

  db, e := connect()
  if e != nil {
    fmt.Println(e)
    return
  }
  defer db.Close()

  stmt, e := db.Prepare("select * from tb_student where id=? and age=?")

  if e != nil {
    fmt.Println(e)
    return
  }

  var res1 = student{}
  stmt.QueryRow("E001",27).Scan(&res1.id, &res1.name, &res1.age, &res1.grade)
  var res2 = student{}
  stmt.QueryRow("W001",28).Scan(&res2.id, &res2.name, &res2.age, &res2.grade)

  fmt.Println(res1)
  // fmt.Println(res1.id)
  fmt.Println(res2)


}

func sqlQueryRow(){
  var db, e = connect()

  if e != nil {
    fmt.Println(e)
    return
  }

  defer db.Close()

  var result = student{}
  var id = "E001"

  e = db. // the dot should be here
          QueryRow("select * from tb_student where id=?", id).
          Scan(&result.id, &result.name, &result.age, &result.grade)

  if e != nil {
    fmt.Println(e)
    return
  }
  fmt.Println(result)
}

func sqlQuery(){
  db, e := connect()
  if e != nil {
    fmt.Println(e)
    return
  }
  defer db.Close()

  var age = 27

  rows, e := db.Query("select * from tb_student where age = ?", age)
  if e != nil {
    fmt.Println(e)
    return
  }
  defer rows.Close()

  var result []student
  for rows.Next(){
    var each = student{}
    var e = rows.Scan(&each.id, &each.name, &each.age, &each.grade)

    if e != nil {
      fmt.Println(e)
      return
    }
    result = append(result, each)
  }


  if rows.Err() != nil {
    fmt.Println(rows.Err())
    return
  }

  for _,each:=range result{
    fmt.Println(each.name)
  }
}

func connect() (*sql.DB, error){
  db, e := sql.Open("mysql", "root:Strikeboy007@tcp(127.0.0.1:3306)/golang_intro")
  if e != nil {
    return nil, e
  }
  return db, nil
}

var data = []Student{
  {"E001", "Ethan", 21},
  {"W001", "Wick", 22},
  {"B001", "Blood", 23},
  {"B002", "Borne", 24},
}

func user(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "application/json")
  // if r.Method == "GET"{
  if r.Method == "POST"{

    var id = r.FormValue("id")
    var result []byte
    var e error

    for _, each:=range data{
      if each.ID == id{
        result, e = json.Marshal(each)

        if e != nil {
          // fmt.Println(e)
          http.Error(w, e.Error(), http.StatusInternalServerError)
          return
        }

        w.Write(result)
        return
      }
    }
    http.Error(w, "User not found", http.StatusBadRequest)
  }
  http.Error(w, "", http.StatusBadRequest)
}

func users(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "application/json")
  if r.Method == "GET"{
    var result, e = json.Marshal(data)

    if e != nil{
      http.Error(w, e.Error(), http.StatusInternalServerError)
      return
    }

    w.Write(result)
    return
  }
  http.Error(w, "", http.StatusBadRequest)
}

type Student struct{
  ID  string  `bson:"id"` //bson used to customize field name for mongoDB (or noSQL in general?)
  Name string `bson:"name"`
  Grade int   `bson:"Grade"`
}

type student struct{
  id string
  name string
  age int
  grade int
}

type User struct{
  Fullname string `json:"Name"` // tag to mapping json data to property struct
  Age int
}

func index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "HEWWO WELKAM BEK")
}

