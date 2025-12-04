package main

import(
  // "fmt"
  // "log"
  // "math/rand"
  "path/filepath"
  "os"
  // "time"
  // "crypto/md5"
)

// const totalFile int = 3000
// const contentLength int = 5000
// var tempPath = filepath.Join()
var TempPath = filepath.Join(os.Getenv("TEMP"), "pipeline")

// func randomString(length int) string{
//   randomizer := rand.New(rand.NewSource(time.Now().Unix()))
//   letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

//   b := make([]rune, length)
//   for i:= range b{
//     b[i] = letters[randomizer.Intn(len(letters))]
//   }
//   return string(b)
// }

// func generateFiles(){
//   os.RemoveAll(tempPath)
//   os.MkdirAll(tempPath, os.ModePerm)

//   for i:=0;i<totalFile;i++{
//     filename := filepath.Join(tempPath, fmt.Sprintf("file-%d.txt",i))
//     // os.Remove(filename)
//     content := randomString(contentLength)
//     err := os.WriteFile(filename, []byte(content), os.ModePerm)
//     if err != nil {
//       log.Println("Error writing file", filename)
//     }

//     if i%100 == 0 && i>0{
//       log.Println(i, "file created")
//     }
//   }
//   log.Printf("%d of total files created", totalFile)
//   // fmt.Println(tempPath)
// }


// func proceed(){
//   // read file, create md5 hash, rename file to file-md5hash.txt
//   counterTotal := 0
//   counterRenamed := 0

//   // walk() is combing all file and directory from a certain directory
//   err := filepath.Walk(tempPath, func(path string, info os.FileInfo, err error) error{
//     if err != nil{
//       return err
//     }
//     // fmt.Println(path)
//     // if it is a sub dir, return immediately
//     if info.IsDir(){
//       return nil
//     }

//     counterTotal++

//     // read file
//     buf, err := os.ReadFile(path)
//     if err != nil{
//       return err
//     }

//     // sum
//     sum := fmt.Sprintf("%x", md5.Sum(buf))

//     // rename file
//     destinationPath := filepath.Join(tempPath, fmt.Sprintf("file-%s.txt", sum))
//     err = os.Rename(path, destinationPath)
//     if err != nil {
//       return err
//     }

//     counterRenamed++
//     return nil
//   })

//   if err!=nil{
//     log.Println("ERROR:", err)
//   }
//   log.Printf("%d/%d files renamed", counterRenamed, counterTotal)
// }

type FileInfo struct{
  FilePath string
  Content []byte
  Sum     string
  IsRenamed bool
}