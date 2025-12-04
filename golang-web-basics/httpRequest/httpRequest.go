
package main

import(
  "fmt"
  "net/http"
  "encoding/json"
  "bytes"
  "net/url"
  // "io"
)

func main(){
  var users, e1 = fetchUsers()
  if e1 != nil {
    fmt.Println(e1)
    return
  }

  for _,each:= range users{
    fmt.Printf("ID: %s\t Name: %s\t Grade: %d\n", each.ID, each.Name, each.Grade)
  }

  var user1, e2 = fetchUser("E001")

  if e2!=nil{
    fmt.Println(user1)
    fmt.Println("Error",e2)
    return
  }
  // fmt.Printf("ID: %s\t Name: %s\t Grade: %d\n", user1.ID, user1.Name, user1.Grade)
  fmt.Println(user1)
}

func fetchUser(ID string)(Student, error){
  var e error
  var client = &http.Client{}
  var data Student

  var param = url.Values{}
  param.Set("id", ID)
  var payload = bytes.NewBufferString(param.Encode())

  request, e := http.NewRequest("POST", baseUrl+"/user", payload)
  // btw, this should be get and baseUrl should just contain user id since it's only fetch something,
  //  it's just tutorial to show how the form passed I guess
  if e != nil {
    return data, e
  }

  request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

  response, e := client.Do(request)
  if e != nil {
    return data, e
  }

  defer response.Body.Close()

  // bodyBytes, _ := io.ReadAll(response.Body)
  // fmt.Println("body",string(bodyBytes))
  // if you uncomment above codes, you'll get empty data since there's a concept of consuming response.Body
  e = json.NewDecoder(response.Body).Decode(&data)
  if e != nil {
    return data, e
  }
  return data, nil
}

func fetchUsers()([]Student, error){
  var e error
  var client = &http.Client{}
  var data []Student

  request, e := http.NewRequest("GET", baseUrl+"/users", nil) //(httpMethod, urlRequest, formRequest)
  if e != nil{
    return nil, e
  }

  response, e := client.Do(request)
  if e != nil {
    return nil, e
  }
  defer response.Body.Close()

  e = json.NewDecoder(response.Body).Decode(&data)
  if e != nil {
    return nil, e
  }

  return data, nil
}

var baseUrl = "http://localhost:8080"

type Student struct{
  ID  string
  Name string
  Grade int
}